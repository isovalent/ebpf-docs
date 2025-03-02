# This is a MkDocs hook and markdown extension that allow for the usage
# of markdown links in code blocks.

from markdown import Extension
from markdown.treeprocessors import Treeprocessor
from markdown.util import AMP_SUBSTITUTE

import xml.etree.ElementTree as etree

import logging
import posixpath

from typing import Iterator

from mkdocs import utils
from mkdocs.utils import _removesuffix
from mkdocs.structure.files import File, Files

from urllib.parse import unquote as urlunquote
from urllib.parse import urlsplit, urlunsplit

log = logging.getLogger(__name__)

# Define a custom markdown extension
class LinkInCodeExtension(Extension):
    # Take the LinksInCodeTreeprocessor as an argument
    def __init__(self, lict):
        self.lict = lict
        super().__init__()

    # Take the markdown instance, add it to the treeprocessor, and register it with
    # the markdown instance.
    def extendMarkdown(self, md):
        # syntax highlighting uses 30, we want to run after
        self.lict.md = md
        md.treeprocessors.register(self.lict, 'links_in_code', 40)

# This is our treeprocessor that will allow us to modify the HTML tree that has
# already been generated. Pygments has already highlighted the code blocks.
class LinksInCodeTreeprocessor(Treeprocessor):
    def __init__(self) -> None:
        self.links_to_anchors: dict[File, dict[str, str]] = {}

    # This method is called on every markdown document, with the HTML root element
    # as the argument. We can modify the tree in place.
    def run(self, root):
        # The `pymdownx.highlight` extension that highlights the code using Pygments
        # stores the output HTML in the `htmlStash`. Which is an external list of HTML.
        # The HTML is re-inserted back at the very end. A placeholder <p> element is
        # inserted with a marker and index.
        # 
        # We want to find these placeholders, find and then modify the HTML stored
        # in the `htmlStash`.
        blocks = root.iter('p')
        for elem in blocks:
            # Does this element have the placeholder marker?
            if elem.text[:1] == '\u0002':
                # Parse the index from the marker, and get the code highlighted html
                end = elem.text.find('\u0003')
                idx = int(elem.text[len('\u0002wzxhzdk:'):end])
                html = self.md.htmlStash.rawHtmlBlocks[idx]

                # Parse the html into a tree we can work with
                try:
                    codeRoot = etree.fromstring(html)
                except etree.ParseError as e:
                    continue

                self.convert_links_in_code(codeRoot)

                # Serialize the modified tree back to HTML string and store it in the `htmlStash`
                self.md.htmlStash.rawHtmlBlocks[idx] = etree.tostring(codeRoot, encoding='unicode', method='html')

    def convert_links_in_code(self, codeRoot):
        # Inside the <code> element, is a flat list of spans with classes to color
        # the different parts of the text.
        codes = codeRoot.iter('code')
        for spans in codes:
            self.convert_links_in_spans(spans)
    
    def convert_links_in_spans(self, spans):
        # Use a `while` loop to iterate over the spans, as we may remove
        # spans from the list. And iterate backwards to avoid messing up the
        # index.
        i = len(spans)
        while i > 0:
            i -= 1

            if spans[i].tag != 'span':
                continue

            # If a span contains more spans, recursively call this method.
            # This happens when lines are highlighted for example, with the `hl_lines` option.
            if len(spans[i]) > 0:
                self.convert_links_in_spans(spans[i])

            # Now look for the markdown link syntax: [text](url)
            # It will be chunked into separate tokens
            # '[' {text} '](' {url} ')'
            #
            # Find the '[' to start a possible match
            if spans[i].text is None:
                continue

            ii = spans[i].text.find('[')
            if ii == -1:
                continue

            # First check if we can find the full link in this one span.
            # This happens when its placed in a comment or macro.
            ji = spans[i].text.find('](', ii)
            ki = spans[i].text.find(')', ji)
            if ji != -1 and ki != -1:
                # Take the text between '[' and '](' as the link text
                link = spans[i].text[ji+2:ki]

                # Create a new <a> element with the href attribute set to the link
                aElem = etree.Element('a', attrib={'href': self.path_to_url(link)})

                # Create a new <span> element with the same class as the original span
                # And add it to the <a> element
                textElem = etree.Element('span', attrib={'class': spans[i].get('class')})
                textElem.text = spans[i].text[ii+1:ji]
                aElem.insert(0, textElem)

                # Make a new span with the same calss as the original span
                # Add the text after the ')' to it.
                postElem = etree.Element('span', attrib={'class': spans[i].get('class')})
                postElem.text = spans[i].text[ki+1:]

                # Move tail from original to post element, to preserve newlines at correct places
                postElem.tail = spans[i].tail
                spans[i].tail = None

                # Insert the link, then post element after the original span
                spans.insert(i+1, aElem)
                spans.insert(i+2, postElem)

                # Truncate text in original to just before the '['
                spans[i].text = spans[i].text[:ii]
                
                # We found a match in a single span, we might find another one.
                # Change index so we process postElem next time.
                i += 3
                continue

            # Use `found` to let the inner look break out of this outer one
            found = False
            # Loop over the rest of the spans until we find a match, or not.
            for j in range(i+1, len(spans)):
                if found:
                    break

                # Find the span that starts with '](', if we do, link text will be
                # between `i` and `j`
                if spans[j].text is None:
                    continue

                ji = spans[j].text.find('](')
                if ji == -1:
                    continue

                # If the tail contains a newline, we are at the end of the line
                # A link can't span multiple lines, so we break out of the loop
                if spans[j].tail is not None and spans[j].tail.find('\n') != -1:
                    break

                # We found the '](', now we have to find the ')'

                # The rest of the text is the link
                link = spans[j].text[ji+2:]
                # Remove the '](' and the text after it from the span
                spans[j].text = spans[j].text[:ji]

                # Loop from `j+1` to the end of the spans
                for k in range(j+1, len(spans)):
                    # If we find a span that starts with ')', we have a match for a full
                    # link pattern. At this point `link` should contain the full href.
                    ki = spans[k].text.find(')')
                    if ki == -1:
                        # If we have not yet found the ')', add the text of the current span
                        # to the link.
                        link += spans[k].text

                        # If the tail contains a newline, we are at the end of the line
                        # A link can't span multiple lines, so we break out of the loop
                        if spans[k].tail is not None and spans[k].tail.find('\n') != -1:
                            Found = True
                            break

                        continue

                    # We have a match, break out of the outer loop after breaking the
                    # inner one.
                    found = True

                    # Remove any text after the `[`
                    spans[i].text = spans[i].text[:ii]

                    # Add the text up to ')' to the link.
                    link += spans[k].text[:ki]           
                    #  Remove the text up to and including the ')' from the text of the span at `k`
                    spans[k].text = spans[k].text[ki+1:]

                    # Call `path_to_url` to convert the path to a markdown file into
                    # a proper URL following the MkDocs config. Normally this is done
                    # by a markdown extension built into MkDocs, but it runs at prio 5.
                    # While code highlighting runs at prio 30. Since we have to run
                    # after the code highlighting, we have to do this manually.
                    # We simply copied the relevant logic from the MkDocs source code.
                    aElem = etree.Element('a', attrib={'href': self.path_to_url(link)})

                    # Now, copy the link text (tokens between `i` and `j`) to the
                    # new `<a>` element.
                    for l in range(j-i-1):
                        aElem.insert(l, spans[i+l+1])

                    # Remove these from the spans list, by repeatedly removing the
                    # i+1'th element.
                    for l in range(k-i-1):
                        spans.remove(spans[i+1])

                    # Insert the new `<a>` element right after the `i`'th element.
                    spans.insert(i+1, aElem)

                    break


    # Copyright © 2014-present, Tom Christie. All rights reserved.
    def path_to_url(self, url: str) -> str:
        scheme, netloc, path, query, anchor = urlsplit(url)

        absolute_link = None
        warning_level, warning = 0, ''

        # Ignore URLs unless they are a relative link to a source file.
        if scheme or netloc:  # External link.
            return url
        elif url.startswith(('/', '\\')):  # Absolute link.
            absolute_link = self.config.validation.links.absolute_links
            if absolute_link is not _AbsoluteLinksValidationValue.RELATIVE_TO_DOCS:
                warning_level = absolute_link
                warning = f"Doc file '{self.file.src_uri}' contains an absolute link '{url}', it was left as is."
        elif AMP_SUBSTITUTE in url:  # AMP_SUBSTITUTE is used internally by Markdown only for email.
            return url
        elif not path:  # Self-link containing only query or anchor.
            if anchor:
                # Register that the page links to itself with an anchor.
                self.links_to_anchors.setdefault(self.file, {}).setdefault(anchor, url)
            return url

        path = urlunquote(path)
        # Determine the filepath of the target.
        possible_target_uris = self._possible_target_uris(
            self.file, path, self.config.use_directory_urls
        )

        if warning:
            # For absolute path (already has a warning), the primary lookup path should be preserved as a tip option.
            target_uri = url
            target_file = None
        else:
            # Validate that the target exists in files collection.
            target_uri = next(possible_target_uris)
            target_file = self.files.get_file_from_path(target_uri)

        if target_file is None and not warning:
            # Primary lookup path had no match, definitely produce a warning, just choose which one.
            if not posixpath.splitext(path)[-1] and absolute_link is None:
                # No '.' in the last part of a path indicates path does not point to a file.
                warning_level = self.config.validation.links.unrecognized_links
                warning = (
                    f"Doc file '{self.file.src_uri}' contains an unrecognized relative link '{url}', "
                    f"it was left as is."
                )
            else:
                target = f" '{target_uri}'" if target_uri != url.lstrip('/') else ""
                warning_level = self.config.validation.links.not_found
                warning = (
                    f"Doc file '{self.file.src_uri}' contains a link '{url}', "
                    f"but the target{target} is not found among documentation files."
                )

        if warning:
            if self.file.inclusion.is_excluded():
                warning_level = min(logging.INFO, warning_level)

            # There was no match, so try to guess what other file could've been intended.
            if warning_level > logging.DEBUG:
                suggest_url = ''
                for path in possible_target_uris:
                    if self.files.get_file_from_path(path) is not None:
                        if anchor and path == self.file.src_uri:
                            path = ''
                        elif absolute_link is _AbsoluteLinksValidationValue.RELATIVE_TO_DOCS:
                            path = '/' + path
                        else:
                            path = utils.get_relative_url(path, self.file.src_uri)
                        suggest_url = urlunsplit(('', '', path, query, anchor))
                        break
                else:
                    if '@' in url and '.' in url and '/' not in url:
                        suggest_url = f'mailto:{url}'
                if suggest_url:
                    warning += f" Did you mean '{suggest_url}'?"
            log.log(warning_level, warning)
            return url

        assert target_uri is not None
        assert target_file is not None

        if anchor:
            # Register that this page links to the target file with an anchor.
            self.links_to_anchors.setdefault(target_file, {}).setdefault(anchor, url)

        if target_file.inclusion.is_excluded():
            if self.file.inclusion.is_excluded():
                warning_level = logging.DEBUG
            else:
                warning_level = min(logging.INFO, self.config.validation.links.not_found)
            warning = (
                f"Doc file '{self.file.src_uri}' contains a link to "
                f"'{target_uri}' which is excluded from the built site."
            )
            log.log(warning_level, warning)
        path = utils.get_relative_url(target_file.url, self.file.url)
        return urlunsplit(('', '', path, query, anchor))

    # Copyright © 2014-present, Tom Christie. All rights reserved.
    @classmethod
    def _target_uri(cls, src_path: str, dest_path: str) -> str:
        return posixpath.normpath(
            posixpath.join(posixpath.dirname(src_path), dest_path).lstrip('/')
        )

    # Copyright © 2014-present, Tom Christie. All rights reserved.
    @classmethod
    def _possible_target_uris(
        cls, file: File, path: str, use_directory_urls: bool, suggest_absolute: bool = False
    ) -> Iterator[str]:
        """First yields the resolved file uri for the link, then proceeds to yield guesses for possible mistakes."""
        target_uri = cls._target_uri(file.src_uri, path)
        yield target_uri

        if posixpath.normpath(path) == '.':
            # Explicitly link to current file.
            yield file.src_uri
            return
        tried = {target_uri}

        prefixes = [target_uri, cls._target_uri(file.url, path)]
        if prefixes[0] == prefixes[1]:
            prefixes.pop()

        suffixes: list[Callable[[str], str]] = []
        if use_directory_urls:
            suffixes.append(lambda p: p)
        if not posixpath.splitext(target_uri)[-1]:
            suffixes.append(lambda p: posixpath.join(p, 'index.md'))
            suffixes.append(lambda p: posixpath.join(p, 'README.md'))
        if (
            not target_uri.endswith('.')
            and not path.endswith('.md')
            and (use_directory_urls or not path.endswith('/'))
        ):
            suffixes.append(lambda p: _removesuffix(p, '.html') + '.md')

        for pref in prefixes:
            for suf in suffixes:
                guess = posixpath.normpath(suf(pref))
                if guess not in tried and not guess.startswith('../'):
                    yield guess
                    tried.add(guess)

# Register a global tree processor so we can modify it from hook callbacks
# while it is being used by the markdown extension.
lict = LinksInCodeTreeprocessor()

# This is called once by MkDocs when config is determined.
# Add our custom markdown extension to the list of markdown extensions.
def on_config(config, **kwargs):
    config.markdown_extensions.append(LinkInCodeExtension(lict))

# This hook is called right before markdown is actually rendered.
# Pass the current file, all files, and config to the markdown extension.
# It will use these to convert links.
def on_page_markdown(markdown, page, config, files, **kwargs):
    lict.file = page.file
    lict.files = files
    lict.config = config
