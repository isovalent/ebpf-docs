package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

var projectroot = flag.String("project-root", "", "Root of the project")
var failFast = flag.Bool("fail-fast", false, "Fail on the first file with misspelled words")

func main() {
	flag.Parse()

	err := checkDir(*projectroot + "/out")
	if err != nil {
		if errors.Is(err, errMisspelled) {
			fmt.Println("Misspelled words found. Please fix them. Or add them to the dictionary (.aspell.en.pws)")
			os.Exit(1)
		}
	}
}

func checkDir(path string) error {
	entry, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	foundMisspelling := false

	for _, e := range entry {
		subPath := path + "/" + e.Name()
		if e.IsDir() {
			if err = checkDir(subPath); err != nil {
				if errors.Is(err, errMisspelled) {
					foundMisspelling = true
					if *failFast {
						break
					}
				} else {
					return err
				}
			}
		} else {
			if err = checkFile(subPath); err != nil {
				if errors.Is(err, errMisspelled) {
					foundMisspelling = true
					if *failFast {
						break
					}
				} else {
					return err
				}
			}
		}
	}

	if foundMisspelling {
		return errMisspelled
	}

	return nil
}

var errMisspelled = errors.New("misspelled")

func checkFile(path string) error {
	if path[len(path)-5:] != ".html" {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	cmd := aspellCmd()
	inPipe, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	// Enter aspell's terse mode. In this mode aspell will only output misspelled words and suggestions and
	// skip printing '*' for each correct word.
	_, err = inPipe.Write([]byte("!\n"))
	if err != nil {
		return err
	}

	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	outScanner := bufio.NewScanner(outPipe)
	outScanner.Split(bufio.ScanLines)

	err = cmd.Start()
	if err != nil {
		return err
	}

	root, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	var misspellings []misspelling

	inGenerated := false
	inHead := false
	visitNode(root, func(n *html.Node) action {
		if inHead {
			p := n.Parent
			for {
				if p.Type == html.ElementNode && p.Data == "head" {
					break
				}

				if p.Parent == nil {
					inHead = false
					break
				}

				p = p.Parent
			}
		}

		switch n.Type {
		case html.DocumentNode, html.DoctypeNode, html.RawNode:
			return next
		case html.ErrorNode:
			return skipSubtree
		case html.CommentNode:
			if strings.HasPrefix(n.Data, " [/HELPER_FUNC_DEF]") ||
				strings.HasPrefix(n.Data, " [/MTU_TABLE]") {
				inGenerated = false
				return next
			}
			if strings.HasPrefix(n.Data, " [HELPER_FUNC_DEF]") ||
				strings.HasPrefix(n.Data, " [MTU_TABLE]") {
				inGenerated = true
			}
		case html.ElementNode:
			switch n.Data {
			case
				// Ignore scripts and style since they don't contain text we own
				"script", "style",
				// Ignore the navigation, every element has its own page anyway
				"nav",
				// Ignore code/quotes in code tags
				"code",
				// Ignore anything inside the custom <nospell> tag
				"nospell":
				return skipSubtree
			case "head":
				inHead = true
			}

			return next
		case html.TextNode:
			if inGenerated {
				return next
			}

			for _, line := range strings.Split(n.Data, "\n") {
				trimmed := strings.TrimSpace(line)
				if trimmed == "" {
					return next
				}

				// Special case, if in <head></head> we use '' instead of `` to mark code / keywords.
				// So remove anything between single quotes.
				if inHead {
					for {
						start := strings.Index(trimmed, "'")
						if start == -1 {
							break
						}

						end := strings.Index(trimmed[start+1:], "'")
						if end == -1 {
							break
						}

						trimmed = trimmed[:start] + trimmed[start+end+2:]
					}
				}

				// If the line starts with a special character, add a space so aspell doesn't take it
				// as a command
				linePadded := false
				if strings.ContainsAny(trimmed[:1], "*&@#~+-!%^$") {
					trimmed = " " + trimmed
					linePadded = true
				}

				_, err = inPipe.Write([]byte(trimmed + "\n"))
				if err != nil {
					slog.Error("Failed to write to aspell", "err", err)
					return skipSubtree
				}

				for {
					if !outScanner.Scan() {
						break
					}

					line := outScanner.Text()
					if line == "" {
						break
					}

					// We shouldn't get these in terse mode, but just in case
					if strings.HasPrefix(line, "*") {
						continue
					}

					if !strings.HasPrefix(line, "&") {
						continue
					}

					parts := strings.SplitN(line, " ", 5)
					misspelled := parts[1]
					offsetInLine, _ := strconv.Atoi(strings.TrimRight(parts[3], ":"))
					if linePadded {
						offsetInLine -= 1
					}
					suggestions := parts[4]
					misspellings = append(misspellings, misspelling{
						file:        path,
						misspelled:  misspelled,
						context:     trimmed,
						offset:      offsetInLine,
						suggestions: strings.Split(suggestions, ", "),
					})
				}
			}
		default:
			return next
		}

		return next
	})

	// Close the pipe to signal the end of input
	err = inPipe.Close()
	if err != nil {
		return err
	}

	// Drain the output
	for outScanner.Scan() {
	}

	// Wait for the command to finish
	err = cmd.Wait()

	if len(misspellings) > 0 {
		for _, m := range misspellings {
			fmt.Printf("Misspelled '%s' found in %s\n", m.misspelled, m.file)
			fmt.Printf("Did you mean one of: %s?\n", strings.Join(m.suggestions, ", "))
			possibleLocations := findPossibleLocations(m.file, m.context)
			if (len(possibleLocations)) > 0 {
				fmt.Printf("Possible locations in markdown:\n")

				for _, loc := range possibleLocations {
					noRoot := strings.TrimPrefix(strings.TrimPrefix(loc.file, *projectroot), "/")
					fmt.Printf("  %s:%d:%d\n", noRoot, loc.line, loc.column+m.offset)
				}
			}
			fmt.Println()
		}

		return errMisspelled
	}

	return nil
}

type fileLocation struct {
	file   string
	line   int
	column int
}

func findPossibleLocations(path string, context string) []fileLocation {
	mdPath := strings.Replace(strings.Replace(path, "/out/", "/docs/", 1), "/index.html", ".md", 1)
	mdFile, err := os.Open(mdPath)
	if err != nil {
		mdPath = strings.Replace(strings.Replace(path, "/out/", "/docs/", 1), ".html", ".md", 1)
		mdFile, err = os.Open(mdPath)
		if err != nil {
			return nil
		}
	}
	defer mdFile.Close()

	contentsBytes, err := io.ReadAll(mdFile)
	if err != nil {
		return nil
	}
	contents := string(contentsBytes)

	possibleLocations := make([]fileLocation, 0)

	off := 0
	for {
		i := strings.Index(contents[off:], context)
		if i == -1 {
			break
		}

		off += i + 1

		possibleLocations = append(possibleLocations, fileLocation{
			file:   mdPath,
			line:   strings.Count(contents[:off], "\n") + 1,
			column: off - strings.LastIndex(contents[:off], "\n") - 1,
		})
	}

	return possibleLocations
}

type misspelling struct {
	file        string
	misspelled  string
	context     string
	offset      int
	suggestions []string
}

type action int

const (
	next action = iota
	skipSubtree
)

func visitNode(n *html.Node, cb func(*html.Node) action) {
	if act := cb(n); act == skipSubtree {
		return
	}

	if n.FirstChild != nil {
		visitNode(n.FirstChild, cb)
	}

	if n.NextSibling != nil {
		visitNode(n.NextSibling, cb)
	}
}

func aspellCmd() *exec.Cmd {
	return exec.Command(
		"aspell",
		"-a",
		"--lang=en",
		"--home-dir="+*projectroot,
	)
}
