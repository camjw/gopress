package scripts

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func GetGitDiffs(basebranch string) []byte {
	cmdName := "git"
	cmdArgs := []string{"diff", "--name-only", basebranch, "HEAD"}
	cmdResponse, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git diff command: ", err)
		os.Exit(1)
	}
	return cmdResponse
}

func RunCypressTests(specsToRun []string) {
	specPath := strings.Join(specsToRun, ",")
	cmd := exec.Command("npx", "cypress", "run", "--spec", specPath)

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running cypress: ", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running cypress: ", err)
		return
	}
}

func CheckRegexesAgainstDiffs(diffs []byte, regexes []string) bool {
	for _, expression := range regexes {
		match, err := regexp.Match(expression, diffs)
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
