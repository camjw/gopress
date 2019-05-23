package main

import (
	"fmt"
	"os"

	parser "github.com/camjw/gopress/internal/parser"
	scripts "github.com/camjw/gopress/internal/scripts"
)

func main() {
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
