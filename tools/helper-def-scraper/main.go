// This tool scrapes the helper definitions from libbpf headers and places them in designated sections in the helper
// function pages.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

const libBpfGhHelperDefsURL = "https://raw.githubusercontent.com/libbpf/libbpf/master/src/bpf_helper_defs.h"

var (
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
		resp, err := http.Get(libBpfGhHelperDefsURL)
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

func manpageToMarkdown(manpage string) string {
	var sb strings.Builder
	scan := bufio.NewScanner(strings.NewReader(manpage))

	type block struct {
		content []string
		indent  int
		code    bool
		header  bool
	}

	titleSeen := false
	var blocks []block
	curBlock := block{
		indent: 1,
	}

	for scan.Scan() {
		line := scan.Text()

		if !strings.HasPrefix(line, " \t") {
			// Lines that start with a single space an no tab are headers
			if strings.HasPrefix(line, " ") {
				// We always skip the title and the next line
				if !titleSeen {
					titleSeen = true
					scan.Scan()
					continue
				}

				// Write a H3
				if len(curBlock.content) > 0 {
					blocks = append(blocks, curBlock)
					curBlock = block{}
				}

				curBlock.header = true
				curBlock.content = []string{line}

				blocks = append(blocks, curBlock)
				curBlock = block{
					indent: 1,
				}

				continue
			}

			// Ignore blank lines
			if line == "" {
				blocks = append(blocks, curBlock)
				curBlock = block{
					indent: curBlock.indent,
				}
				continue
			}

			panic("unknown line format")
		}

		level := strings.Count(line, "\t")

		if curBlock.indent != level {
			if len(curBlock.content) > 0 {
				blocks = append(blocks, curBlock)
			}
			curBlock = block{
				indent:  level,
				content: nil,
			}
		}

		line = strings.TrimPrefix(line, " ")
		line = strings.ReplaceAll(line, "\t", "")

		if line == "::" {
			curBlock.code = true
			curBlock.indent++
			scan.Scan()
			continue
		}

		curBlock.content = append(curBlock.content, line)
	}
	if len(curBlock.content) > 0 {
		blocks = append(blocks, curBlock)
	}

	for i, block := range blocks {
		if i != 0 {
			sb.WriteString("\n")
		}

		if block.code {
			sb.WriteString("```\n")
		}

		if block.header {
			sb.WriteString("###")
		}

		if !block.code && !block.header {
			sb.WriteString(strings.Repeat("&nbsp;", (block.indent-1)*4))
		}

		for ii, line := range block.content {
			var lb strings.Builder
			isListItem := false
			for i := 0; i < len(line); i++ {
				// "* " at the beginning of the line, is a list
				if i == 0 && line[i] == '*' && i+1 < len(line) && line[i+1] == ' ' {
					lb.WriteByte(line[i])
					lb.WriteByte(line[i+1])
					i++
					isListItem = true
					continue
				}

				// "- " at the beginning of the line, is a list
				if i == 0 && line[i] == '-' && i+1 < len(line) && line[i+1] == ' ' {
					lb.WriteByte(line[i])
					lb.WriteByte(line[i+1])
					i++
					isListItem = true
					continue
				}

				// Cut escaped spaces
				if line[i] == '\\' && i+1 < len(line) && line[i+1] == ' ' {
					i++
					continue
				}

				nextIsAsterisk := i+1 < len(line) && line[i+1] == '*'
				prevIsAsterisk := i > 0 && line[i-1] == '*'

				// Only replace single *, not **
				if line[i] == '*' && (!nextIsAsterisk && !prevIsAsterisk) {
					lb.WriteByte('_')
					continue
				}

				lb.WriteByte(line[i])
			}

			sb.WriteString(lb.String())

			if isListItem {
				sb.WriteString("\n")
				if ii+1 < len(block.content) {
					sb.WriteString(strings.Repeat("&nbsp;", block.indent-1))
				}
			} else {
				if ii+1 < len(block.content) {
					sb.WriteString(" ")
				}
			}
		}

		sb.WriteString("\n")

		if block.code {
			sb.WriteString("```\n")
		}
	}

	return sb.String()
}
