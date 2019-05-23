package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	parser "github.com/camjw/gopress/internal/parser"
)

func getFileChanges(basebranch string) []byte {
	cmdName := "git"
	cmdArgs := []string{"diff", "--name-only", basebranch, "HEAD"}
	cmdResponse, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git diff command: ", err)
		os.Exit(1)
	}
	return cmdResponse
}

func checkRegexesAgainstChanges(changes []byte, regexes []string) bool {
	for _, expression := range regexes {
		match, err := regexp.Match(expression, changes)
		if match {
			return true
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "One of the regexes is malformed: ", expression, "Error ocurred: ", err)
			os.Exit(1)
		}
	}
	return false
}

func runCypressTests(specs []string) {
	cmdName := "npx"
	fmt.Println(specs)
	_, err := exec.Command(cmdName, specs...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running cypress: ", err)
	}
}

func main() {
	config, err := parser.GetConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error parsing the goparser json: ", err)
		os.Exit(1)
	}

	testcases := config.Tests
	fileBytes := getFileChanges(config.Basebranch)

	specsToRun := []string{"cypress", "run", "--specs"}

	for testIdx, _ := range testcases {
		testcase := testcases[testIdx]
		if checkRegexesAgainstChanges(fileBytes, testcase.Regexes) {
			specsToRun = append(specsToRun, config.Directory+testcase.Testfile+config.Extension)
		}
	}

	runCypressTests(specsToRun)
}
