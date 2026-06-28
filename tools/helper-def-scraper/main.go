// This tool scrapes the helper definitions from libbpf headers and places them in designated sections in the helper
// function pages.
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

const libBpfGhHelperDefsURL = "https://raw.githubusercontent.com/libbpf/libbpf/{ref}/src/bpf_helper_defs.h"

var (
	libbpfRef      = flag.String("libbpf-ref", "master", "libbpf ref")
	filePath       = flag.String("file-path", "", "If set, use a file path instead of fetching from the interwebs")
	helperFuncPath = flag.String("helper-path", "", "The path the helper function pages")

	helperRegex = regexp.MustCompile(`static [^\(]+ \*?\(\* const ([^\)]+)\)[^\n]+;`)
)

const (
	helperDefMarkerStart = "<!-- [HELPER_FUNC_DEF] -->\n"
	helperDefMarkerStop  = "<!-- [/HELPER_FUNC_DEF] -->\n"
)

func main() {
	flag.Parse()

	if helperFuncPath == nil || *helperFuncPath == "" {
		fmt.Fprintln(os.Stderr, "Missing flag 'helper-path'")
		return
	}

	var headerFileReader io.Reader
	if filePath != nil && *filePath != "" {
		file, err := os.Open(*filePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: %w", err)
			return
		}
		defer file.Close()

		headerFileReader = file
	} else {
		url := strings.Replace(libBpfGhHelperDefsURL, "{ref}", *libbpfRef, 1)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: %w", err)
			return
		}
		defer resp.Body.Close()

		headerFileReader = resp.Body
	}

	headerFile, err := io.ReadAll(headerFileReader)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: %w", err)
		return
	}

	definitions := make(map[string]helperDef)
	var (
		curHelper helperDef
		inComment bool
	)

	for _, line := range strings.Split(string(headerFile), "\n") {
		if strings.TrimSpace(line) == "/*" {
			curHelper = helperDef{
				Description: "",
			}
			inComment = true
			continue
		}

		if strings.TrimSpace(line) == "*/" {
			inComment = false
			continue
		}

		if inComment {
			line = strings.TrimPrefix(line, " *")
			curHelper.Description += line + "\n"
		}

		if helperRegex.MatchString(line) {
			matches := helperRegex.FindStringSubmatch(line)
			curHelper.Definition = matches[0]
			curHelper.Name = matches[1]
			definitions[curHelper.Name] = curHelper
		}
	}

	for name, def := range definitions {
		func() {
			mapPath := path.Join(*helperFuncPath, name+".md")
			fmt.Printf("Opening '%s'\n", mapPath)

			mapFile, err := os.OpenFile(mapPath, os.O_RDWR, 0644)
			if err != nil {
				fmt.Printf("Skipping '%s' due to error: %s\n", name, err.Error())
				return
			}

			defer mapFile.Close()

			fileContents, err := io.ReadAll(mapFile)
			if err != nil {
				fmt.Printf("Skipping '%s' due to error: %s\n", name, err.Error())
				return
			}
			fileStr := string(fileContents)

			startIdx := strings.Index(fileStr, helperDefMarkerStart)
			if startIdx == -1 {
				fmt.Printf("Skipping '%s', missing ref start marker\n", name)
				return
			}

			stopIdx := strings.Index(fileStr, helperDefMarkerStop)
			if stopIdx == -1 {
				fmt.Printf("Skipping '%s', missing ref stop marker\n", name)
				return
			}

			var newFile strings.Builder
			// Write everything before the marker
			newFile.WriteString(fileStr[:startIdx])
			newFile.WriteString(helperDefMarkerStart)

			newFile.WriteString(renderHelperFuncDef(def))

			newFile.WriteString(helperDefMarkerStop)
			newFile.WriteString(fileStr[stopIdx+len(helperDefMarkerStop):])

			_, err = mapFile.Seek(0, 0)
			if err != nil {
				fmt.Printf("Skipping '%s', seek error\n", name)
				return
			}

			err = mapFile.Truncate(0)
			if err != nil {
				fmt.Printf("Skipping '%s', truncate error\n", name)
				return
			}

			_, err = io.Copy(mapFile, strings.NewReader(newFile.String()))
			if err != nil {
				fmt.Printf("Skipping '%s', copy error\n", name)
				return
			}

		}()
	}
}

type helperDef struct {
	Name        string
	Definition  string
	Description string
}

func renderHelperFuncDef(def helperDef) string {
	var sb strings.Builder

	// sb.WriteString("<!-- src\n")
	// sb.WriteString(def.Description)
	// sb.WriteString("\n-->")

	sb.WriteString(manpageToMarkdown(def.Description))
	sb.WriteString("\n")
	sb.WriteString("`#!c ")
	sb.WriteString(def.Definition)
	sb.WriteString("`\n")

	return sb.String()
}

// manpageToMarkdown converts the helper comment text from libbpf's
// manpage-like format into markdown used in docs pages.
func manpageToMarkdown(manpage string) string {
	// We intentionally skip the first line because in helper comments it is the helper name title.
	// the second line is a blank space so we skip it as well.
	lines := strings.Split(manpage, "\n")[2:]
	// we remove the initial space from each line, so that we only have tabs
	for i, s := range lines {
		lines[i] = strings.TrimPrefix(s, " ")
	}

	parser := newBlockParser(blockParserArgs{
		lines:  lines,
		debug:  false,
		output: os.Stdout,
	})
	return renderMarkdownBlocks(parser.generateBlocks())
}

func renderMarkdownBlocks(blocks []markdownBlock) string {
	var sb strings.Builder
	for i, b := range blocks {
		if i > 0 {
			// we don't want extra newlines between list items
			if blocks[i-1].isList && b.isList {
				sb.WriteByte('\n')
			} else {
				sb.WriteString("\n\n")
			}
		}
		sb.WriteString(b.text)
	}

	// we need a final newline
	sb.WriteByte('\n')
	return sb.String()
}
