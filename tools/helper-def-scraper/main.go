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

const libBpfGhHelperDefsURL = "https://raw.githubusercontent.com/libbpf/libbpf/master/src/bpf_helper_defs.h"

var (
	filePath       = flag.String("file-path", "", "If set, use a file path instead of fetching from the interwebs")
	helperFuncPath = flag.String("helper-path", "", "The path the the helper function pages")

	helperRegex = regexp.MustCompile(`static [^\(]+ \*?\(\*([^\)]+)\)[^\n]+;`)
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

	lines := strings.Split(def.Description, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		line = strings.Replace(line, "**", "`", -1)
		line = strings.Replace(line, "*", "`", -1)

		if strings.HasPrefix(line, " \t") {
			sb.WriteString(strings.TrimSpace(line) + "\n")
			continue
		} else if strings.TrimSpace(line) == "" {
			sb.WriteString("\n")
			continue
		}

		// Ignore the header of the same name as the function.
		if strings.TrimSpace(line) == def.Name {
			// Skip the blank line following the first header
			i++
			continue
		}

		// Format header
		sb.WriteString("\n**" + strings.TrimSpace(line) + "**\n")
	}

	sb.WriteString("`#!c ")
	sb.WriteString(def.Definition)
	sb.WriteString("`\n")

	return sb.String()
}
