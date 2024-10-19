package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

type timelineInfo struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	Commit string `yaml:"commit"`
}

const (
	timelinePath = "docs/linux/timeline/index.md"
)

func dumpTimeline(timeline map[string][]timelineInfo) error {
	filePath := path.Join(*projectroot, timelinePath)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	_, err = file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("seek: %w", err)
	}

	err = file.Truncate(0)
	if err != nil {
		return fmt.Errorf("truncate: %w", err)
	}

	var sb strings.Builder

	sb.WriteString(`---
title: eBPF Timeline
description: This page lists all eBPF features added in the Linux Kernel ordered by tag number.
hide: toc
---
`)

	// Print the tags ordered. These are SemVer strings but to avoid external modules we compare only the Major and Minor
	var tags []string
	for tag := range timeline {
		tags = append(tags, tag)
	}
	sort.Slice(tags, func(i, j int) bool {
		// We need to remove the initial `v` from the tag.
		v1Parts := strings.Split(tags[i][1:], ".")
		v2Parts := strings.Split(tags[j][1:], ".")

		v1Major, _ := strconv.Atoi(v1Parts[0])
		v1Minor, _ := strconv.Atoi(v1Parts[1])
		v2Major, _ := strconv.Atoi(v2Parts[0])
		v2Minor, _ := strconv.Atoi(v2Parts[1])

		// Compare Major versions
		if v1Major != v2Major {
			return v1Major < v2Major
		}

		// In our case it will never happen to have 2 tags with the same Major and Minor, but in any case we return true if they are equal
		return v1Minor <= v2Minor
	})

	for _, tag := range tags {
		fmt.Fprintf(&sb, "\n## :octicons-tag-24: %s\n\n", tag)
		for _, feature := range timeline[tag] {
			fmt.Fprintf(&sb, "* **%s** [%s](https://github.com/torvalds/linux/commit/%s) (%s)\n", feature.Name, feature.Commit[:7], feature.Commit, feature.Type)
		}
	}

	_, err = io.Copy(file, strings.NewReader(sb.String()))
	if err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	return nil
}

func generateTimeline(file dataFile) error {
	// Generate the timeline from the data file
	timeline := make(map[string][]timelineInfo)
	for _, group := range file {
		for _, feature := range group.Features {
			timeline[feature.Version] = append(timeline[feature.Version], timelineInfo{Name: feature.Name, Type: group.Name, Commit: feature.Commit})
		}
	}
	return dumpTimeline(timeline)
}
