// This tool looks for markers in doc pages that request the insertion of a version tag for a given "feature"
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
)

var projectroot = flag.String("project-root", "", "Root of the project")

const (
	programsDir = "docs/linux/program-type"
)

var mtuTableStartRegex = regexp.MustCompile(`<!-- \[MTU_TABLE\] -->`)

const mtuTableMarkerStop = "<!-- [/MTU_TABLE] -->"

func main() {
	flag.Parse()

	if projectroot == nil || *projectroot == "" {
		fmt.Fprintln(os.Stderr, "Missing 'project-root' flag")
		return
	}

	if err := mainE(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func mainE() error {
	entries, err := os.ReadDir(path.Join(*projectroot, programsDir))
	if err != nil {
		return fmt.Errorf("Read dir: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := path.Join(*projectroot, programsDir, entry.Name())
		err = processFile(filePath)
		if err != nil {
			return fmt.Errorf("Process file: %w", err)
		}
	}

	return nil
}

func processFile(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("Open file: %w", err)
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Read all: %w", err)
	}

	// If the file doesn't contain any markers, don't do anything
	if !mtuTableStartRegex.Match(contents) {
		return nil
	}

	contentsStr := string(contents)
	var sb strings.Builder

	for {
		loc := mtuTableStartRegex.FindStringSubmatchIndex(contentsStr)
		// No more matches
		if loc == nil {
			// Write the remaining content
			sb.WriteString(contentsStr)
			break
		}

		stopIndex := strings.Index(contentsStr, mtuTableMarkerStop)
		if stopIndex == -1 {
			// Write the remaining content
			sb.WriteString(contentsStr)
			break
		}

		// Write content including start tag
		sb.WriteString(contentsStr[:loc[1]])

		printTableMatrix(&sb)

		// Write stop marker
		sb.WriteString(contentsStr[stopIndex : stopIndex+len(mtuTableMarkerStop)])

		// Shrink off processed portion
		contentsStr = contentsStr[stopIndex+len(mtuTableMarkerStop):]
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("seek: %w", err)
	}

	err = file.Truncate(0)
	if err != nil {
		return fmt.Errorf("truncate: %w", err)
	}

	_, err = io.Copy(file, strings.NewReader(sb.String()))
	if err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	return nil
}
