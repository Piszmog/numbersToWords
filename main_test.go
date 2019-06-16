package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	tempWriteFile, err := ioutil.TempFile(os.TempDir(), "testInput-")
	if err != nil {
		t.Fatal("failed to create temp file for test", err)
	}
	//
	// defers are last-in-first-out
	//
	defer os.Remove(tempWriteFile.Name())
	_, err = tempWriteFile.WriteString("1\n")
	if err != nil {
		t.Fatal("failed to write line", err)
	}
	_, err = tempWriteFile.WriteString("2\n")
	if err != nil {
		t.Fatal("failed to write line", err)
	}
	_, err = tempWriteFile.WriteString("3\n")
	if err != nil {
		t.Fatal("failed to write line", err)
	}
	closeFile(tempWriteFile)
	file, err := openFile(tempWriteFile.Name())
	if err != nil {
		t.Fatalf("failed to open temp file at %s: %+v", tempWriteFile.Name(), err)
	}
	lineChannel := make(chan string, 3)
	readFile(file, lineChannel)
	line := <-lineChannel
	if line != "1" {
		t.Fatalf("expected first line to be '1' but got '%s'", line)
	}
	line = <-lineChannel
	if line != "2" {
		t.Fatalf("expected first line to be '2' but got '%s'", line)
	}
	line = <-lineChannel
	if line != "3" {
		t.Fatalf("expected first line to be '3' but got '%s'", line)
	}
}
