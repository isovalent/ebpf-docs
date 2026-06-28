package main

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

var (
	// This is used to match bold text and put the placeholder in its place
	boldToken = regexp.MustCompile(`\*\*[^*\n]+\*\*`)
	// This is used to match helper function calls
	helperCallToken = regexp.MustCompile(`\*\*([^*\n]+)\*\*\s+\(\)`)
	// Match italic inline markers like *flags*, *addr*, *task->mm*, *percpu map*.
	italicToken = regexp.MustCompile(`\*([A-Za-z0-9_.>/\-]+(?: [A-Za-z0-9_.>/\-]+)*)\*`)
	// - start with `\t` and followed by optional \t or spaces.
	// - followed by `*`
	// - after `*` we could have \t or spaces
	listItemRe = regexp.MustCompile(`^\t[ \t]*\*[ \t]+`)
)

type markdownBlock struct {
	isList bool
	text   string
}

type blockParser struct {
	lines  []string
	cursor int
	blocks []markdownBlock
	debug  bool
	output io.Writer
}

type blockParserArgs struct {
	lines  []string
	debug  bool
	output io.Writer
}

func newBlockParser(args blockParserArgs) *blockParser {
	const defaultBlockCapacity = 16 // this seems a reasonable default
	return &blockParser{
		lines:  args.lines,
		cursor: 0,
		blocks: make([]markdownBlock, 0, defaultBlockCapacity),
		debug:  args.debug,
		output: args.output,
	}
}

func (p *blockParser) debugBlockMsg(block *markdownBlock) {
	if !p.debug {
		return
	}

	if p.output == nil {
		panic("output is nil, cannot print debug message")
	}

	_, _ = fmt.Fprintf(p.output, "Block n%d:\n{\n%s\n}\n", len(p.blocks)+1, block.text)
}

func (p *blockParser) appendBlock(block markdownBlock) {
	p.debugBlockMsg(&block)
	p.blocks = append(p.blocks, block)
}

func cleanInlineMarkdown(text string) string {
	// remove space escape sequences
	text = strings.ReplaceAll(text, "\\ ", "")

	// normalize helper names from `**bpf_xxx** ()` to `**bpf_xxx**()`.
	text = helperCallToken.ReplaceAllString(text, "**$1**()")

	// finds all `**...**` and replaces them with placeholders
	// so that we don't have conflicts with italic replacements
	boldPlaceholders := make([]string, 0, 8)
	text = boldToken.ReplaceAllStringFunc(text, func(s string) string {
		placeholder := fmt.Sprintf("__BOLD_BLOCK_%d__", len(boldPlaceholders))
		boldPlaceholders = append(boldPlaceholders, s)
		return placeholder
	})

	// convert `*something*` into markdown format `_something_`
	text = italicToken.ReplaceAllString(text, "_${1}_")

	// restore bold text from placeholders
	for i, block := range boldPlaceholders {
		placeholder := fmt.Sprintf("__BOLD_BLOCK_%d__", i)
		text = strings.ReplaceAll(text, placeholder, block)
	}

	// correct some cases like `(_tuple_ **->ipv4**)` into ``_tuple_**->ipv4**``
	text = strings.ReplaceAll(text, "_ **->", "_**->")

	return text
}

func trimIndent(line string) string {
	return strings.TrimLeft(line, "\t ")
}

func computeIndent(line string) int {
	indent := 0
	for i := 0; i < len(line); i++ {
		if line[i] == '\t' || line[i] == ' ' {
			indent++
		} else {
			break
		}
	}
	return indent
}

func (p *blockParser) parseCodeBlock() {
	const codePrefix = "\t\t"
	var sb strings.Builder
	sb.WriteString("```\n")
	// we increment the cursor immediately by 2
	// because we skip the first "::" and the empty line after it.
	i := p.cursor + 2
	for ; i < len(p.lines); i++ {
		// When we have less than 2 tab the code block is completed
		if p.lines[i] != "" && !strings.HasPrefix(p.lines[i], codePrefix) {
			break
		}
		// here we also append blank lines because we need to preserve them in code blocks.
		sb.WriteString(strings.TrimPrefix(p.lines[i], codePrefix) + "\n")
	}
	p.cursor = i
	// the last line of a code block is always an empty line so we remove it.
	p.appendBlock(markdownBlock{text: strings.TrimSuffix(sb.String(), "\n") + "```"})
}

func (p *blockParser) parseBlockUntilCondition(prefix string, terminateFunc func(line string) bool) markdownBlock {
	var sb strings.Builder
	if prefix != "" {
		sb.WriteString(prefix)
	}

	i := p.cursor
	for ; i < len(p.lines); i++ {
		if terminateFunc(p.lines[i]) {
			break
		}

		// skip empty lines if they are not a termination condition
		if p.lines[i] == "" {
			continue
		}

		if sb.Len() > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(trimIndent(p.lines[i]))
	}
	p.cursor = i
	return markdownBlock{text: cleanInlineMarkdown(sb.String())}
}

func (p *blockParser) parseTextBlock() {
	indent := computeIndent(p.lines[p.cursor])
	block := p.parseBlockUntilCondition("", func(line string) bool {
		// termination conditions:
		// - blank line
		// - different indent from the first line
		return line == "" || computeIndent(line) != indent
	})
	p.appendBlock(block)
}

func (p *blockParser) parsePseudoListItemBlock() {
	// The current line is something like `**BPF_MTU_CHK_SEGS**`
	// We want to obtain `* **BPF_MTU_CHK_SEGS**:`
	prefix := "* " + trimIndent(p.lines[p.cursor]) + ":"
	indent := computeIndent(p.lines[p.cursor])
	p.cursor += 1
	block := p.parseBlockUntilCondition(prefix, func(line string) bool {
		// termination conditions:
		// - lower indent than the first line excluding blank lines
		return line != "" && computeIndent(line) <= indent
	})
	p.appendBlock(block)
}

func (p *blockParser) parseListItemBlock() {
	prefix := trimIndent(p.lines[p.cursor])
	indent := computeIndent(p.lines[p.cursor])
	p.cursor += 1
	block := p.parseBlockUntilCondition(prefix, func(line string) bool {
		// termination conditions:
		// - lower indent than the first line excluding blank lines
		return line != "" && computeIndent(line) <= indent
	})
	block.isList = true
	p.appendBlock(block)
}

func (p *blockParser) parseHeaderBlock() {
	p.appendBlock(markdownBlock{
		text: "### " + strings.TrimLeft(p.lines[p.cursor], " "),
	})
	p.cursor++
}

func (p *blockParser) isPseudoListItem() bool {
	/* With "pseudo-list" we mean blocks that should behave like lists
	 * but they don't start with a single `*`.
	 *
	 * Example:
	 *
	 * 	**BTF_F_COMPACT**
	 * 		no formatting around type information
	 * 	**BTF_F_NONAME**
	 * 		no struct/union member names/types
	 * 	**BTF_F_PTR_RAW**
	 * 		show raw (unobfuscated) pointer values;
	 * 		equivalent to printk specifier %px.
	 * 	**BTF_F_ZERO**
	 * 		show zero-valued struct/union members; they
	 * 		are not displayed by default
	 *
	 * Example:
	 * 	**sizeof**\ (*tuple*\ **->ipv4**)
	 * 		Look for an IPv4 socket.
	 * 	**sizeof**\ (*tuple*\ **->ipv6**)
	 * 		Look for an IPv6 socket.
	 *
	 * For now we just check for initial `\t**` and the current line shouldn't be the last one.
	 */

	currLine := p.lines[p.cursor]
	if !strings.HasPrefix(currLine, "\t**") || p.cursor+1 >= len(p.lines) {
		return false
	}
	// The next line must be indented more than the current line.
	return computeIndent(p.lines[p.cursor+1]) > computeIndent(currLine)
}

func (p *blockParser) isCodeBlock() bool {
	// Every code block starts with a line containing `\t::`
	return strings.HasPrefix(p.lines[p.cursor], "\t::")
}

func (p *blockParser) isListItem() bool {
	return listItemRe.MatchString(p.lines[p.cursor])
}

func (p *blockParser) isTextBlock() bool {
	return strings.HasPrefix(p.lines[p.cursor], "\t")
}

func (p *blockParser) isWhitespace() bool {
	return p.lines[p.cursor] == ""
}

func (p *blockParser) generateBlocks() []markdownBlock {
	for p.cursor < len(p.lines) {
		// The order counts, so we need to check in this order.
		if p.isWhitespace() {
			// In case of empty lines between blocks we skip them.
			p.cursor++
			// todo!: add exceptions for certain cases...
		} else if p.isCodeBlock() {
			p.parseCodeBlock()
		} else if p.isPseudoListItem() {
			p.parsePseudoListItemBlock()
		} else if p.isListItem() {
			p.parseListItemBlock()
		} else if p.isTextBlock() {
			p.parseTextBlock()
		} else {
			// If there are no `\t` at the beginning,
			// the only case we are aware of is a header.
			p.parseHeaderBlock()
		}
	}
	return p.blocks
}
