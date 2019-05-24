package main

import "testing"

func TestVersionFlag(t *testing.T) {

	setupInputs([]string{"--version"}, nil)
	main()
}
