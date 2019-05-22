package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
)

type testRegex struct {
	Filename string   `json:"filename"`
	Regexes  []string `json:"regexes"`
}

func getFileChanges() []byte {
	cmdName := "git"
	cmdArgs := []string{"diff", "--name-only", "HEAD^", "HEAD"}
	cmdResponse, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git diff command: ", err)
		os.Exit(1)
	}
	return cmdResponse
}

func getTestRegexes(filename string) []testRegex {
	regexesFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error reading the features json: ", err)
		os.Exit(1)
	}
	var testRegexes []testRegex
	err = json.Unmarshal(regexesFile, &testRegexes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "The test regexes json could not be read: ", err)
		os.Exit(1)
	}
	return testRegexes
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

	testRegexes := getTestRegexes("./cypress/features.json")
	fileBytes := getFileChanges()

	specsToRun := []string{"cypress", "run", "--specs"}

	for regexIdx, _ := range testRegexes {
		testRegex := testRegexes[regexIdx]
		if checkRegexesAgainstChanges(fileBytes, testRegex.Regexes) {
			specsToRun = append(specsToRun, testRegex.Filename)
		}
	}

	runCypressTests(specsToRun)
}
