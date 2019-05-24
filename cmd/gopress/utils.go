package main

import (
	"flag"
	"os"
)

var stdin *os.File

func setupInputs(args []string, file *os.File) {

	a := os.Args[1:]
	if args != nil {
		a = args
	}

	flag.CommandLine.Parse(a)

	if file != nil {
		stdin = file
	} else {
		stdin = os.Stdin
	}
}
