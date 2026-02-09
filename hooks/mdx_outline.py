'''
Outline extension for Python-Markdown
=====================================

Wraps the document logical sections (as implied by h1-h6 headings).

By default, the wrapper element is a section tag having a class attribute
"sectionN", where N is the header level being wrapped. Also, the header
attributes get moved to the wrapper.


Usage
-----

    >>> import markdown
    >>> src = """
    ... # 1
    ... Section 1
    ... ## 1.1
    ... Subsection 1.1
    ... ## 1.2
    ... Subsection 1.2
    ... ### 1.2.1
    ... Hey 1.2.1 Special section
    ... ### 1.2.2
    ... #### 1.2.2.1
    ... # 2
    ... Section 2
    ... """.strip()
    >>> html = markdown.markdown(src, extensions=['outline'])
    >>> print(html)
    <section class="section1"><h1>1</h1>
    <p>Section 1</p>
    <section class="section2"><h2>1.1</h2>
    <p>Subsection 1.1</p>
    </section><section class="section2"><h2>1.2</h2>
    <p>Subsection 1.2</p>
    <section class="section3"><h3>1.2.1</h3>
    <p>Hey 1.2.1 Special section</p>
    </section><section class="section3"><h3>1.2.2</h3>
    <section class="section4"><h4>1.2.2.1</h4>
    </section></section></section></section><section class="section1"><h1>2</h1>
    <p>Section 2</p>
    </section>

Divs instead of sections, custom class names:

    >>> src = """
    ... # Introduction
    ... # Body
    ... ## Subsection
    ... # Bibliography
    ... """.strip()
    >>> html = markdown.markdown(src, extensions=['outline(wrapper_tag=div, wrapper_cls=s%(LEVEL)d)'])
    >>> print(html)
    <div class="s1"><h1>Introduction</h1>
    </div><div class="s1"><h1>Body</h1>
    <div class="s2"><h2>Subsection</h2>
    </div></div><div class="s1"><h1>Bibliography</h1>
    </div>


By default, the header attributes are moved to the wrappers

    >>> src = """
    ... # Introduction {: foo='bar' }
    ... """.strip()
    >>> html = markdown.markdown(src, extensions=['attr_list', 'outline'])
    >>> print(html)
    <section class="section1" foo="bar"><h1>Introduction</h1>
    </section>


Content-specified classes are added to settings wrapper class

    >>> src = """
    ... # Introduction {: class='extraclass' }
    ... """.strip()
    >>> html = markdown.markdown(src, extensions=['attr_list', 'outline'])
    >>> print(html)
    <section class="extraclass section1"><h1>Introduction</h1>
    </section>


Non consecutive headers shouldn't be a problem:
    >>> src="""
    ... # ONE
    ... ### TOO Deep
    ... ## Level 2
    ... # TWO
    ... """.strip()
    >>> html = markdown.markdown(src, extensions=['attr_list', 'outline'])
    >>> print(html)
    <section class="section1"><h1>ONE</h1>
    <section class="section3"><h3>TOO Deep</h3>
    </section><section class="section2"><h2>Level 2</h2>
    </section></section><section class="section1"><h1>TWO</h1>
    </section>


Dependencies
------------

* [Markdown 2.0+](http://www.freewisdom.org/projects/python-markdown/)


Copyright
---------

- 2011, 2012 [The active archives contributors](http://activearchives.org/)
- 2011, 2012 [Michael Murtaugh](http://automatist.org/)
- 2011, 2012, 2017 [Alexandre Leray](http://stdin.fr/)

All rights reserved.

This software is released under the modified BSD License. 
See LICENSE.md for details.


Further credits
---------------

This is a rewrite of the 
[mdx_addsection extension](http://git.constantvzw.org/?p=aa.core.git;a=blob;f=aacore/mdx_addsections.py;h=969e520a42b0018a2c4b74889fecc83a7dd7704a;hb=HEAD) 
we've written for [active archives](http://activearchives.org). The first
version had a bug with non hierarchical heading structures. This is no longer a
problem: a couple of weeks ago, Jesse Dhillon pushed to github a similar plugin
which fixes the problem. Thanks to him, mdx_outline no longer has the problem.


See also
--------

- <https://github.com/jessedhillon/mdx_sections>
- <http://html5doctor.com/outlines/>
'''


import re
from xml.etree import ElementTree

from markdown import Extension
from markdown.treeprocessors import Treeprocessor

__version__ = "1.4.0"


class OutlineProcessor(Treeprocessor):
    def process_nodes(self, node):
        s = []
        pattern = re.compile("^h(\d)")
        wrapper_cls = self.wrapper_cls

        for child in list(node):
            match = pattern.match(child.tag.lower())

            if match:
                depth = int(match.group(1))

                section = ElementTree.SubElement(node, self.wrapper_tag)
                section.append(child)

                if self.move_attrib:
                    for key, value in list(child.attrib.items()):
                        section.set(key, value)
                        del child.attrib[key]

                node.remove(child)

                if "%(LEVEL)d" in self.wrapper_cls:
                    wrapper_cls = self.wrapper_cls % {"LEVEL": depth}

                cls = section.attrib.get("class")
                if cls:
                    section.attrib["class"] = " ".join([cls, wrapper_cls])
                elif wrapper_cls:  # no class attribute if wrapper_cls==''
                    section.attrib["class"] = wrapper_cls

                contained = False

                while s:
                    container, container_depth = s[-1]
                    if depth <= container_depth:
                        s.pop()
                    else:
                        contained = True
                        break

                if contained:
                    container.append(section)
                    node.remove(section)

                s.append((section, depth))

            else:
                if s:
                    container, container_depth = s[-1]
                    container.append(child)
                    node.remove(child)

    def run(self, root):
        self.wrapper_tag = self.config.get("wrapper_tag")[0]
        self.wrapper_cls = self.config.get("wrapper_cls")[0]
        self.move_attrib = self.config.get("move_attrib")[0]

        self.process_nodes(root)
        return root


class OutlineExtension(Extension):
    def __init__(self, *args, **kwargs):
        self.config = {
            "wrapper_tag": ["section", "Tag name to use, default: section"],
            "wrapper_cls": [
                "section%(LEVEL)d",
                "Default CSS class applied to sections",
            ],
            "move_attrib": [True, "Move header attributes to the wrapper"],
        }
        super(OutlineExtension, self).__init__(**kwargs)

    def extendMarkdown(self, md, md_globals=None):
        ext = OutlineProcessor(md)
        ext.config = self.config
        md.treeprocessors.register(ext, "outline", 10)

def on_config(config, **kwargs):
    configs={}
    config.markdown_extensions.append(OutlineExtension(configs))
