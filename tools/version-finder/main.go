// This tool finds the patterns in git repositories and records what the earliest version was where it occurred.
// Such information can then be used to generate documentation about when features were added.
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	_ "embed"

	"gopkg.in/yaml.v3"
)

type patternFile struct {
	Files    []string       `yaml:"files"`
	Tags     []Regexp       `yaml:"tags"`
	Patterns []patternGroup `yaml:"patterns"`
}

type patternGroup struct {
	Name     string    `yaml:"name"`
	Patterns []pattern `yaml:"patterns"`
}

type pattern struct {
	Name    string   `yaml:"name"`
	Regexes []string `yaml:"regexes"`
}

type Regexp struct {
	*regexp.Regexp
}

func (r *Regexp) UnmarshalYAML(value *yaml.Node) (err error) {
	r.Regexp, err = regexp.Compile(value.Value)
	return err
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

//go:embed patterns.yaml
var configFile string

var (
	kernelDir    = flag.String("kernel-dir", "", "Path to the linux kernel directory")
	dataFilePath = flag.String("data-file", "", "Path to the data file output")
)

func main() {
	flag.Parse()
	if kernelDir == nil || *kernelDir == "" {
		fmt.Fprintln(os.Stderr, "Error: --kernel-dir is required")
		return
	}

	if dataFilePath == nil || *dataFilePath == "" {
		fmt.Fprintln(os.Stderr, "Error: --data-file is required")
		return
	}

	var config patternFile
	err := yaml.Unmarshal([]byte(configFile), &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: yaml unmarshal: ", err)
		return
	}

	err = os.Chdir(*kernelDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: chdir: ", err)
		return
	}

	tags, err := gitTags()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: git tags: ", err)
		return
	}

	var matchedTags []string
	for _, tag := range tags {
		for _, pattern := range config.Tags {
			if !pattern.MatchString(tag) {
				continue
			}

			matchedTags = append(matchedTags, tag)
			break
		}
	}

	dataFileWriter, err := os.Create(*dataFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: create file: ", err)
		return
	}

	type versions struct {
		PrevVersion string
		Version     string
		Commit      string
	}
	matches := make(map[string]map[string]versions)

	lastTag := ""

	// Walk over tags from newest to oldest, detect when a feature no longer appears
	for _, tag := range matchedTags {
		fmt.Println("Checkout ", tag)
		err := gitCheckout(tag)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: git checkout: ", err)
			return
		}

		for _, group := range config.Patterns {
			fmt.Println("Greping ", group.Name)
			groupMatches := matches[group.Name]
			if groupMatches == nil {
				groupMatches = make(map[string]versions)
			}

			for _, pattern := range group.Patterns {
				// If we didn't match on the last iteration, don't check it again since we will only be going
				// further back in time. This will save us a lot of useless duplication
				if groupMatches[pattern.Name].Version != lastTag {
					continue
				}

				regexes := pattern.Regexes
				if len(regexes) == 0 {
					regexes = append(regexes, pattern.Name)
				}

				found, err := gitGrep(config.Files, regexes)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error: git grep: ", err)
					return
				}

				if !found {
					versions := groupMatches[pattern.Name]
					versions.PrevVersion = tag
					groupMatches[pattern.Name] = versions
					continue
				}

				groupMatches[pattern.Name] = versions{
					Version: tag,
				}
			}

			matches[group.Name] = groupMatches
		}

		lastTag = tag
	}

	// Loop over all features and bisect to find the exact commit
	for _, patternGroup := range config.Patterns {
		groupMatches := matches[patternGroup.Name]
		for _, pattern := range patternGroup.Patterns {
			versions := groupMatches[pattern.Name]
			if versions.Version == "" || versions.PrevVersion == "" {
				continue
			}

			fmt.Println("Bisecting", versions.PrevVersion, versions.Version, pattern.Name)

			regexes := pattern.Regexes
			if len(regexes) == 0 {
				regexes = append(regexes, pattern.Name)
			}

			commit, err := gitBisect(versions.PrevVersion, versions.Version, config.Files, regexes)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error: git bisect: ", err)
				return
			}
			versions.Commit = commit
			groupMatches[pattern.Name] = versions
		}

		matches[patternGroup.Name] = groupMatches
	}

	err = gitBisectReset()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: git bisect reset: ", err)
		return
	}

	var dataFileContent dataFile

	for _, patternGroup := range config.Patterns {
		groupMatches := matches[patternGroup.Name]
		dataGroup := dataGroup{
			Name: patternGroup.Name,
		}
		for _, feature := range patternGroup.Patterns {
			dataGroup.Features = append(dataGroup.Features, featureVersion{
				Name:    feature.Name,
				Version: groupMatches[feature.Name].Version,
				Commit:  groupMatches[feature.Name].Commit,
			})
		}
		dataFileContent = append(dataFileContent, dataGroup)
	}

	err = yaml.NewEncoder(dataFileWriter).Encode(dataFileContent)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: encode matches: ", err)
		return
	}
}

func gitGrep(files, patterns []string) (bool, error) {
	var args []string

	args = append(args, "grep")

	// Quite mode makes the command throw and error if no matches are found
	args = append(args, "-q")

	// Only trigger on word boundaries to prevent partial matches
	args = append(args, "-w")

	// Patterns are OR-ed by default
	for _, pattern := range patterns {
		args = append(args, "-e", pattern)
	}

	args = append(args, "--")
	args = append(args, files...)

	cmd := exec.Command("git", args...)

	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func gitCheckout(ref string) error {
	cmd := exec.Command("git", "checkout", ref)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// gitTags returns all tags in the git repository from newest to oldest
func gitTags() ([]string, error) {
	var out bytes.Buffer
	cmd := exec.Command("git", "for-each-ref", "--sort", "-creatordate", "--format", "%(refname:short)", "refs/tags")
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return strings.Split(out.String(), "\n"), nil
}

func gitBisect(prevVersion, version string, files, patterns []string) (string, error) {
	cmd := exec.Command("git", "bisect", "start", version, prevVersion)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	var args []string

	// Do an inverted grep, so error if we find any match
	args = append(args, "!", "git", "grep")

	// Quite mode makes the command throw and error if no matches are found
	args = append(args, "-q")

	// Only trigger on word boundaries to prevent partial matches
	args = append(args, "-w")

	// Patterns are OR-ed by default
	for _, pattern := range patterns {
		args = append(args, "-e", pattern)
	}

	args = append(args, "--")
	args = append(args, files...)

	// Run the git grep for every step, once done, HEAD will be the commit to add the feature
	cmd = exec.Command("git", "bisect", "run", "sh", "-c", strings.Join(args, " "))
	if err := cmd.Run(); err != nil {
		return "", err
	}

	cmd = exec.Command("git", "rev-parse", "refs/bisect/bad")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func gitBisectReset() error {
	cmd := exec.Command("git", "bisect", "reset")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
