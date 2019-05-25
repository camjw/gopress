package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	parser "gopress/internal/parser"
	scripts "gopress/internal/scripts"
)

func runGopress(file io.Reader) {
	config, err := parser.GetConfig(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error parsing the goparser json: ", err)
		os.Exit(1)
	}

	var specsToRun []string
	testcases := config.Tests
	fileBytes := scripts.GetGitDiffs(config.Basebranch)

	for testIdx, _ := range testcases {
		testcase := testcases[testIdx]
		if scripts.CheckRegexesAgainstDiffs(fileBytes, testcase.Regexes) {
			specsToRun = append(specsToRun, config.GetFilePath(testcase))
		}
	}

	if len(specsToRun) > 0 {
		scripts.RunCypressTests(specsToRun)
	} else {
		fmt.Println("No specs to run")
	}
}

func loadFile(filename string) (io.Reader, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("There was an error loading the gopress file: " + err.Error())
	}
	return file, nil
}

func handleFlags(versionFlag *bool) {
	if *versionFlag {
		fmt.Println("v.0.0-alpha")
		os.Exit(0)
	}

	file, err := loadFile("./gopress.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	runGopress(file)
}

func main() {
	versionFlag := flag.Bool("version", false, "Check the version of Gopress")

	flag.Parse()

	handleFlags(versionFlag)
}
