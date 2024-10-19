package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
)

var featureTagMarkerStartRegex = regexp.MustCompile(`<!-- \[FEATURE_TAG\]\(([^\)]+)\) -->`)

const (
	programsDir          = "docs/linux/program-type"
	mapDir               = "docs/linux/map-type"
	helperDir            = "docs/linux/helper-function"
	syscallDir           = "docs/linux/syscall"
	kfuncsDir            = "docs/linux/kfuncs"
	featureTagMarkerStop = "<!-- [/FEATURE_TAG] -->"
)

func processFile(path string, featureMap map[string]featureVersion) error {
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
	if !featureTagMarkerStartRegex.Match(contents) {
		return nil
	}

	contentsStr := string(contents)
	var sb strings.Builder

	for {
		loc := featureTagMarkerStartRegex.FindStringSubmatchIndex(contentsStr)
		// No more matches
		if loc == nil {
			// Write the remaining content
			sb.WriteString(contentsStr)
			break
		}

		featureName := contentsStr[loc[2]:loc[3]]

		stopIndex := strings.Index(contentsStr, featureTagMarkerStop)
		if stopIndex == -1 {
			// Write the remaining content
			sb.WriteString(contentsStr)
			break
		}

		// Write content including start tag
		sb.WriteString(contentsStr[:loc[1]])

		// Write tag
		version, found := featureMap[featureName]
		if found {
			fmt.Fprintf(&sb, "\n[:octicons-tag-24: %s](https://github.com/torvalds/linux/commit/%s)\n", version.Version, version.Commit)
		} else {
			sb.WriteString("\n:octicons-tag-24: unknown\n")
		}

		// Write stop marker
		sb.WriteString(contentsStr[stopIndex : stopIndex+len(featureTagMarkerStop)])

		// Shrink off processed portion
		contentsStr = contentsStr[stopIndex+len(featureTagMarkerStop):]
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

func generateTags(file dataFile) error {
	featureMap := make(map[string]featureVersion)
	for _, group := range file {
		for _, feature := range group.Features {
			featureMap[feature.Name] = feature
		}
	}

	for _, dir := range []string{programsDir, helperDir, mapDir, syscallDir, kfuncsDir} {
		entries, err := os.ReadDir(path.Join(*projectroot, dir))
		if err != nil {
			return fmt.Errorf("Read dir: %w", err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			filePath := path.Join(*projectroot, dir, entry.Name())
			err = processFile(filePath, featureMap)
			if err != nil {
				return fmt.Errorf("Process file: %w", err)
			}
		}
	}
	return nil
}
