package main

import (
	"flag"
	"fmt"
	"os"

	parser "github.com/camjw/gopress/internal/parser"
	scripts "github.com/camjw/gopress/internal/scripts"
)

func runGopress() {
	config, err := parser.GetConfig()
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

	scripts.RunCypressTests(specsToRun)
}

func main() {
	versionFlag := flag.Bool("version", false, "Check the version of Gopress")

	flag.Parse()

	if *versionFlag {
		fmt.Println("v.0.0-alpha")
	} else {
		runGopress()
	}
}
