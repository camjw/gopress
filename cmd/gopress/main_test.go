package main

import "testing"

var T = true

func ExampleHandleFlags() {
	handleFlags(&T)
	// Output: v.0.0-alpha
}

func TestLoadFile(t *testing.T) {
	_, err := loadFile("./notafile")

	got := err.Error()
	want := "There was an error loading the gopress file: open ./notafile: no such file or directory"

	if got != want {
		t.Errorf("Got: %s. Want: %s", got, want)
	}
}
