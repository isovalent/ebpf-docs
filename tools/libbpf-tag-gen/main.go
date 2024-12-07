package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	projectroot = flag.String("project-root", "", "Root of the project")
	libbpfRef   = flag.String("libbpf-ref", "master", "libbpf ref")
)

const libbpfMapURL = "https://raw.githubusercontent.com/libbpf/libbpf/{ref}/src/libbpf.map"

const (
	LIBBPF_TAG_START = "<!-- [LIBBPF_TAG] -->"
	LIBBPF_TAG_END   = "<!-- [/LIBBPF_TAG] -->"
)

func main() {
	flag.Parse()
	if *projectroot == "" {
		panic("project-root is required")
	}

	url := strings.Replace(libbpfMapURL, "{ref}", *libbpfRef, 1)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to download libbpf.map: %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	funcToTag, err := parseLibbpfMap(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse libbpf.map: %v\n", err)
		os.Exit(1)
	}

	dirPath := path.Join(*projectroot, "docs", "ebpf-library", "libbpf", "userspace")
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read directory: %v\n", err)
		os.Exit(1)
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}
		if !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		entrypath := path.Join([]string{*projectroot, "docs", "ebpf-library", "libbpf", "userspace", entry.Name()}...)
		file, err := os.OpenFile(entrypath, os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
			continue
		}

		fileContents, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
			continue
		}

		funcName := strings.TrimSuffix(entry.Name(), ".md")
		tag, ok := funcToTag[funcName]
		if !ok {
			fmt.Fprintf(os.Stderr, "Function %s not found in libbpf.map\n", funcName)
			continue
		}

		fileStr := string(fileContents)
		startIndex := strings.Index(fileStr, LIBBPF_TAG_START)
		endIndex := strings.Index(fileStr, LIBBPF_TAG_END)

		if startIndex == -1 || endIndex == -1 {
			fmt.Fprintf(os.Stderr, "Skipping, can not find tag markers in file '%s'\n", entry.Name())
			continue
		}

		var newFile strings.Builder
		// Write everything before the marker
		newFile.WriteString(fileStr[:startIndex])
		newFile.WriteString(LIBBPF_TAG_START)
		fmt.Fprintf(&newFile, "\n[:octicons-tag-24: %s](https://github.com/libbpf/libbpf/releases/tag/v%s)\n", tag, tag)
		newFile.WriteString(LIBBPF_TAG_END)
		newFile.WriteString(fileStr[endIndex+len(LIBBPF_TAG_END):])

		_, err = file.Seek(0, 0)
		if err != nil {
			panic(err)
		}

		err = file.Truncate(0)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(file, strings.NewReader(newFile.String()))
		if err != nil {
			panic(err)
		}
		file.Close()
	}
}

func parseLibbpfMap(body io.Reader) (map[string]string, error) {
	scan := bufio.NewScanner(body)
	funcToTag := make(map[string]string)

	for scan.Scan() {
		line := scan.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "LIBBPF_") {
			if err := parseBlock(scan, funcToTag); err != nil {
				return nil, err
			}
		}
	}

	return funcToTag, nil
}

func parseBlock(scan *bufio.Scanner, funcToTag map[string]string) error {
	line := scan.Text()
	bareLine := strings.TrimSpace(line)
	fields := strings.Fields(bareLine)
	if len(fields) == 0 {
		return nil
	}

	tag := strings.TrimPrefix(fields[0], "LIBBPF_")

	for scan.Scan() {
		line := scan.Text()
		bareLine := strings.TrimSpace(line)
		if strings.HasPrefix(bareLine, "}") {
			break
		}

		if strings.HasSuffix(bareLine, ":") {
			continue
		}

		funcToTag[bareLine[:len(bareLine)-1]] = tag
	}

	return nil
}
