// This tool generates helper function reference sections.
// On program type pages it will make a reference of all helper functions that particular program type supports.
// On helper function pages it will make a reference to program types and maps to list with which it works.
// On map type pages it will make a reference to which helper calls work for the given map.

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"golang.org/x/exp/slices"
	yaml "gopkg.in/yaml.v3"
)

var projectroot = flag.String("project-root", "", "Root of the project")

const (
	dataFilePath = "data/helpers-functions.yaml"
	programsDir  = "docs/linux/program-type"
	mapDir       = "docs/linux/map-type"
	helperDir    = "docs/linux/helper-function"
)

const (
	progRefMarkerStart       = "<!-- [PROG_HELPER_FUNC_REF] -->\n"
	progRefMarkerStop        = "<!-- [/PROG_HELPER_FUNC_REF] -->\n"
	mapRefMarkerStart        = "<!-- [MAP_HELPER_FUNC_REF] -->\n"
	mapRefMarkerStop         = "<!-- [/MAP_HELPER_FUNC_REF] -->\n"
	helperProgRefMarkerStart = "<!-- [HELPER_FUNC_PROG_REF] -->\n"
	helperProgRefMarkerStop  = "<!-- [/HELPER_FUNC_PROG_REF] -->\n"
	helperMapRefMarkerStart  = "<!-- [HELPER_FUNC_MAP_REF] -->\n"
	helperMapRefMarkerStop   = "<!-- [/HELPER_FUNC_MAP_REF] -->\n"
)

type helperFuncGroup []helperDef

type helperDef struct {
	Name               string   `yaml:"name"`
	GroupName          string   `yaml:"group"`
	KConfig            []string `yaml:"kconfig"`
	Capabilities       []string `yaml:"cap"`
	AttachType         []string `yaml:"attach_type"`
	ProgramTrampoline  bool     `yaml:"prog_trampoline"`
	NoSecurityLockdown bool     `yaml:"no_security_lockdown"`
}

type helperFuncDataFile struct {
	Groups   map[string]helperFuncGroup `yaml:"groups"`
	Programs map[string]helperFuncGroup `yaml:"programs"`
	Maps     map[string]helperFuncGroup `yaml:"maps"`
}

func (df *helperFuncDataFile) flatten(group helperFuncGroup) helperFuncGroup {
	for _, member := range group {
		if member.GroupName != "" {
			subGroup := df.Groups[member.GroupName]
			subGroup = df.flatten(subGroup)
			group = append(group, subGroup...)
		}
	}

	for i := len(group) - 1; i >= 0; i-- {
		if group[i].GroupName != "" {
			group = slices.Delete(group, i, i+1)
		}
	}

	return group
}

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
	dataFile, err := parseDataFile()
	if err != nil {
		return err
	}

	err = renderProgramPages(dataFile)
	if err != nil {
		return err
	}

	err = renderMapPages(dataFile)
	if err != nil {
		return err
	}

	err = renderHelperFuncPages(dataFile)
	if err != nil {
		return err
	}

	return nil
}

func renderProgramPages(dataFile *helperFuncDataFile) error {
	for programType := range dataFile.Programs {
		fmt.Printf("Prog type '%s'\n", programType)
		func() {
			progPath := path.Join(*projectroot, programsDir, programType+".md")
			fmt.Printf("Opening '%s'\n", progPath)

			programFile, err := os.OpenFile(progPath, os.O_RDWR, 0644)
			if err != nil {
				fmt.Printf("Skipping '%s' due to error: %s\n", programType, err.Error())
				return
			}

			defer programFile.Close()

			fileContents, err := io.ReadAll(programFile)
			if err != nil {
				fmt.Printf("Skipping '%s' due to error: %s\n", programType, err.Error())
				return
			}
			fileStr := string(fileContents)

			startIdx := strings.Index(fileStr, progRefMarkerStart)
			if startIdx == -1 {
				fmt.Printf("Skipping '%s', missing ref start marker\n", programType)
				return
			}

			stopIdx := strings.Index(fileStr, progRefMarkerStop)
			if stopIdx == -1 {
				fmt.Printf("Skipping '%s', missing ref stop marker\n", programType)
				return
			}

			var newFile strings.Builder
			// Write everything before the marker
			newFile.WriteString(fileStr[:startIdx])
			newFile.WriteString(progRefMarkerStart)

			newFile.WriteString(renderProgramHelperFuncReference(dataFile, programType))

			newFile.WriteString(progRefMarkerStop)
			newFile.WriteString(fileStr[stopIdx+len(progRefMarkerStop):])

			_, err = programFile.Seek(0, 0)
			if err != nil {
				fmt.Printf("Skipping '%s', seek error\n", programType)
				return
			}

			err = programFile.Truncate(0)
			if err != nil {
				fmt.Printf("Skipping '%s', truncate error\n", programType)
				return
			}

			_, err = io.Copy(programFile, strings.NewReader(newFile.String()))
			if err != nil {
				fmt.Printf("Skipping '%s', copy error\n", programType)
				return
			}
		}()
	}

	return nil
}

func renderMapPages(dataFile *helperFuncDataFile) error {
	for mapType := range dataFile.Maps {
		fmt.Printf("Map type '%s'\n", mapType)
		func() {
			mapPath := path.Join(*projectroot, mapDir, mapType+".md")
			fmt.Printf("Opening '%s'\n", mapPath)

			mapFile, err := os.OpenFile(mapPath, os.O_RDWR, 0644)
			if err != nil {
				fmt.Printf("Skipping '%s' due to error: %s\n", mapType, err.Error())
				return
			}

			defer mapFile.Close()

			fileContents, err := io.ReadAll(mapFile)
			if err != nil {
				fmt.Printf("Skipping '%s' due to error: %s\n", mapType, err.Error())
				return
			}
			fileStr := string(fileContents)

			startIdx := strings.Index(fileStr, mapRefMarkerStart)
			if startIdx == -1 {
				fmt.Printf("Skipping '%s', missing ref start marker\n", mapType)
				return
			}

			stopIdx := strings.Index(fileStr, mapRefMarkerStop)
			if stopIdx == -1 {
				fmt.Printf("Skipping '%s', missing ref stop marker\n", mapType)
				return
			}

			var newFile strings.Builder
			// Write everything before the marker
			newFile.WriteString(fileStr[:startIdx])
			newFile.WriteString(mapRefMarkerStart)

			newFile.WriteString(renderMapHelperFuncReference(dataFile, mapType))

			newFile.WriteString(mapRefMarkerStop)
			newFile.WriteString(fileStr[stopIdx+len(mapRefMarkerStop):])

			_, err = mapFile.Seek(0, 0)
			if err != nil {
				fmt.Printf("Skipping '%s', seek error\n", mapType)
				return
			}

			err = mapFile.Truncate(0)
			if err != nil {
				fmt.Printf("Skipping '%s', truncate error\n", mapType)
				return
			}

			_, err = io.Copy(mapFile, strings.NewReader(newFile.String()))
			if err != nil {
				fmt.Printf("Skipping '%s', copy error\n", mapType)
				return
			}
		}()
	}

	return nil
}

func renderHelperFuncPages(dataFile *helperFuncDataFile) error {
	progTypesPerFunc := make(map[string][]string)
	for progType := range dataFile.Programs {
		group := dataFile.flatten(dataFile.Programs[progType])
		for _, helperDef := range group {
			progTypesPerFunc[helperDef.Name] = append(progTypesPerFunc[helperDef.Name], progType)
		}
	}

	for helper, programTypes := range progTypesPerFunc {
		fmt.Printf("Helper func '%s'\n", helper)
		func() {
			mapPath := path.Join(*projectroot, helperDir, helper+".md")
			fmt.Printf("Opening '%s'\n", mapPath)

			mapFile, err := os.OpenFile(mapPath, os.O_RDWR, 0644)
			if err != nil {
				fmt.Printf("Skipping '%s' due to error: %s\n", helper, err.Error())
				return
			}

			defer mapFile.Close()

			fileContents, err := io.ReadAll(mapFile)
			if err != nil {
				fmt.Printf("Skipping '%s' due to error: %s\n", helper, err.Error())
				return
			}
			fileStr := string(fileContents)

			startIdx := strings.Index(fileStr, helperProgRefMarkerStart)
			if startIdx == -1 {
				fmt.Printf("Skipping '%s', missing ref start marker\n", helper)
				return
			}

			stopIdx := strings.Index(fileStr, helperProgRefMarkerStop)
			if stopIdx == -1 {
				fmt.Printf("Skipping '%s', missing ref stop marker\n", helper)
				return
			}

			var newFile strings.Builder
			// Write everything before the marker
			newFile.WriteString(fileStr[:startIdx])
			newFile.WriteString(helperProgRefMarkerStart)

			newFile.WriteString(renderHelperFuncProgReference(programTypes))

			newFile.WriteString(helperProgRefMarkerStop)
			newFile.WriteString(fileStr[stopIdx+len(helperProgRefMarkerStop):])

			_, err = mapFile.Seek(0, 0)
			if err != nil {
				fmt.Printf("Skipping '%s', seek error\n", helper)
				return
			}

			err = mapFile.Truncate(0)
			if err != nil {
				fmt.Printf("Skipping '%s', truncate error\n", helper)
				return
			}

			_, err = io.Copy(mapFile, strings.NewReader(newFile.String()))
			if err != nil {
				fmt.Printf("Skipping '%s', copy error\n", helper)
				return
			}
		}()
	}

	mapTypesPerFunc := make(map[string][]string)
	for mapType := range dataFile.Maps {
		group := dataFile.flatten(dataFile.Maps[mapType])
		for _, helperDef := range group {
			mapTypesPerFunc[helperDef.Name] = append(mapTypesPerFunc[helperDef.Name], mapType)
		}
	}

	for helper, mapTypes := range mapTypesPerFunc {
		fmt.Printf("Helper func '%s'\n", helper)
		func() {
			mapPath := path.Join(*projectroot, helperDir, helper+".md")
			fmt.Printf("Opening '%s'\n", mapPath)

			mapFile, err := os.OpenFile(mapPath, os.O_RDWR, 0644)
			if err != nil {
				fmt.Printf("Skipping '%s' due to error: %s\n", helper, err.Error())
				return
			}

			defer mapFile.Close()

			fileContents, err := io.ReadAll(mapFile)
			if err != nil {
				fmt.Printf("Skipping '%s' due to error: %s\n", helper, err.Error())
				return
			}
			fileStr := string(fileContents)

			startIdx := strings.Index(fileStr, helperMapRefMarkerStart)
			if startIdx == -1 {
				fmt.Printf("Skipping '%s', missing ref start marker\n", helper)
				return
			}

			stopIdx := strings.Index(fileStr, helperMapRefMarkerStop)
			if stopIdx == -1 {
				fmt.Printf("Skipping '%s', missing ref stop marker\n", helper)
				return
			}

			var newFile strings.Builder
			// Write everything before the marker
			newFile.WriteString(fileStr[:startIdx])
			newFile.WriteString(helperMapRefMarkerStart)

			newFile.WriteString(renderHelperFuncMapReference(mapTypes))

			newFile.WriteString(helperMapRefMarkerStop)
			newFile.WriteString(fileStr[stopIdx+len(helperMapRefMarkerStop):])

			_, err = mapFile.Seek(0, 0)
			if err != nil {
				fmt.Printf("Skipping '%s', seek error\n", helper)
				return
			}

			err = mapFile.Truncate(0)
			if err != nil {
				fmt.Printf("Skipping '%s', truncate error\n", helper)
				return
			}

			_, err = io.Copy(mapFile, strings.NewReader(newFile.String()))
			if err != nil {
				fmt.Printf("Skipping '%s', copy error\n", helper)
				return
			}
		}()
	}

	return nil
}

func parseDataFile() (*helperFuncDataFile, error) {
	dataFile, err := os.Open(path.Join(*projectroot, dataFilePath))
	if err != nil {
		return nil, fmt.Errorf("error opening data file: %w", err)
	}
	defer dataFile.Close()

	var helperData helperFuncDataFile
	err = yaml.NewDecoder(dataFile).Decode(&helperData)
	if err != nil {
		return nil, fmt.Errorf("error decoding data file: %w", err)
	}

	return &helperData, nil
}

// The reference of helper functions placed on the program type page
func renderProgramHelperFuncReference(file *helperFuncDataFile, progType string) string {
	group := file.flatten(file.Programs[progType])

	var sb strings.Builder

	sb.WriteString("??? abstract \"Supported helper functions\"\n")
	for _, item := range group {
		sb.WriteString(fmt.Sprintf("    * [%s](../helper-function/%s.md)\n", item.Name, item.Name))
	}

	return sb.String()
}

// The reference of helper functions placed on the program type page
func renderMapHelperFuncReference(file *helperFuncDataFile, progType string) string {
	group := file.flatten(file.Maps[progType])

	var sb strings.Builder

	for _, item := range group {
		sb.WriteString(fmt.Sprintf(" * [%s](../helper-function/%s.md)\n", item.Name, item.Name))
	}

	return sb.String()
}

func renderHelperFuncProgReference(programTypes []string) string {
	var sb strings.Builder

	for _, item := range programTypes {
		sb.WriteString(fmt.Sprintf(" * [%s](../program-type/%s.md)\n", item, item))
	}

	return sb.String()
}

func renderHelperFuncMapReference(mapTypes []string) string {
	var sb strings.Builder

	for _, item := range mapTypes {
		sb.WriteString(fmt.Sprintf(" * [%s](../map-type/%s.md)\n", item, item))
	}

	return sb.String()
}
