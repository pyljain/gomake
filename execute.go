package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func executeTarget(gomakeFileContent, target string) {
	start := strings.Index(gomakeFileContent, target+":")
	if start == -1 {
		fmt.Printf("Target %s not found in Makefile\n", target)
		os.Exit(-1)
	}

	currentLocation := start
	lineContents := bytes.NewBufferString("")
	var targetLines []string

	for currentLocation < len(gomakeFileContent) {
		if gomakeFileContent[currentLocation] == '\n' {
			if strings.Trim(lineContents.String(), "\n") == "" {
				break
			}

			targetLines = append(targetLines, trimString(lineContents.String(), " ", "\n", "\t"))
			lineContents = bytes.NewBufferString("")
		}

		lineContents.WriteByte(gomakeFileContent[currentLocation])

		currentLocation += 1
	}
	targetLines = append(targetLines, trimString(lineContents.String(), " ", "\n", "\t"))

	dependenciesLine := strings.Trim(strings.Split(targetLines[0], ":")[1], " ")
	if len(dependenciesLine) > 1 {
		dependencies := strings.Split(dependenciesLine, " ")
		log.Printf("Dependencies are %+v", dependencies)

		if len(dependencies) > 0 {
			for _, d := range dependencies {
				executeTarget(gomakeFileContent, d)
			}
		}
	}

	err := executeMakeCommands(targetLines[1:])
	if err != nil {
		log.Printf("Error occured when executing commands")
		os.Exit(-1)
	}
}
