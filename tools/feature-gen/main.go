// This tool looks for markers in doc pages that request the insertion of a version tag for a given "feature"
package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var projectroot = flag.String(projectrootFlagName, "", "Root of the project")
var genTimeline = flag.Bool(timelineFlagName, false, "Generate a timeline for eBPF features")
var genTags = flag.Bool(tagsFlagName, false, "Generate tags for eBPF features")

const (
	timelineFlagName    = "timeline"
	tagsFlagName        = "tags"
	projectrootFlagName = "project-root"
	dataFilePath        = "data/feature-versions.yaml"
)

func main() {
	flag.Parse()

	if projectroot == nil || *projectroot == "" {
		fmt.Fprintf(os.Stderr, "missing '%s' flag\n", projectrootFlagName)
		return
	}

	if err := mainE(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func mainE() error {
	dataFile, err := openDataFile()
	if err != nil {
		return fmt.Errorf("open data file: %w", err)
	}

	if *genTimeline {
		return generateTimeline(dataFile)
	} else if *genTags {
		return generateTags(dataFile)
	} else {
		return fmt.Errorf("missing generation flag. Specify '%s' or '%s'", timelineFlagName, tagsFlagName)
	}
}

type dataFile []dataGroup

type dataGroup struct {
	Name     string           `yaml:"name"`
	Features []featureVersion `yaml:"features"`
}

type featureVersion struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Commit  string `yaml:"commit"`
}

func openDataFile() (dataFile, error) {
	file, err := os.Open(path.Join(*projectroot, dataFilePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var df dataFile
	err = yaml.NewDecoder(file).Decode(&df)
	if err != nil {
		return nil, err
	}

	return df, nil
}
