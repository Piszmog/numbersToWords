package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	//
	// Setup flags to read file
	//
	fileLocation := flag.String("f", "", "File containing the numbers to convert to words")
	flag.Parse()
	//
	// Validate input flag
	//
	if len(*fileLocation) == 0 {
		fmt.Println("Missing required flag '-f'")
		flag.Usage()
		return
	}
	//
	// Open the file
	//
	file, err := openFile(*fileLocation)
	if err != nil {
		log.Fatalln(err)
	}
	defer closeFile(file)
	defer printRuntime(time.Now())
	lineChannel := make(chan string, 50)
	go readFile(file, lineChannel)
	handleLines(lineChannel)
}

func openFile(filename string) (*os.File, error) {
	pathToFile, err := filepath.Abs(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get absolute path of %s", filename)
	}
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file %s", filename)
	}
	return file, nil
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(errors.Wrapf(err, "failed to close %s", file.Name()))
	}
}

func readFile(file *os.File, lines chan string) {
	defer close(lines)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
}

func handleLines(lineChannel chan string) {
	for line := range lineChannel {
		if !numberValidator.isStringValid(line) {
			fmt.Printf("Error: Input '%s' is not a valid input.\n", line)
			continue
		}
		//
		// convert to int64
		//
		input, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Printf("failed to convert input '%s' to int64: %+v", line, err)
		} else {
			fmt.Printf("Input: %s\n", line)
			fmt.Printf("Output: %s\n", number(input))
		}
	}
}

func printRuntime(now time.Time) {
	fmt.Printf("\nApplication took %f seconds to complete\n", time.Since(now).Seconds())
}
