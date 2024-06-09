package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	target := os.Args[1]
	if target == "" {
		log.Println("unimplemented")
	}

	gomakeFileContentBytes, err := os.ReadFile("Makefile")
	if err != nil {
		fmt.Printf("No Makefile found in current directory\n")
		os.Exit(-1)
	}

	gomakeFileContent := string(gomakeFileContentBytes)

	executeTarget(gomakeFileContent, target)
}

func trimString(s string, vals ...string) string {
	result := s

	for _, t := range vals {
		result = strings.Trim(result, t)
	}

	return result
}

func executeMakeCommands(commands []string) error {
	for _, cmd := range commands {
		if cmd == "" {
			continue
		}

		c := strings.Split(cmd, " ")
		output, err := exec.Command(c[0], c[1:]...).CombinedOutput()
		if err != nil {
			log.Printf("Error occured running the command %s", cmd)
		}

		fmt.Println(cmd)
		fmt.Print(string(output))
	}

	return nil
}
