package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "git"
	cmdArgs := []string{"diff", "--name-only", "HEAD^", "HEAD"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git diff command: ", err)
		os.Exit(1)
	}
	filenames := string(cmdOut)
	fmt.Println(filenames)
}
