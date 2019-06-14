package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//
	// create regex to ensure only numbers are being converted to int64
	//
	validator, err := regexp.Compile("^[0-9]+$")
	if err != nil {
		log.Fatal("failed to compile validator:", err)
	}
	//
	// create standard input reader and start reading
	//
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a number to convert: ")
		inputString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("\nfailed to read input:", err)
		}
		//
		// remove the '\n' before performing regex or conversion
		//
		inputString = strings.Replace(inputString, "\n", "", -1)
		//
		// validate input
		//
		if !validator.MatchString(inputString) {
			fmt.Printf("Input '%s' is not a valid input. Please try again.\n\n", inputString)
			continue
		}
		//
		// convert to int64
		//
		input, err := strconv.ParseInt(inputString, 10, 64)
		if err != nil {
			log.Fatalf("failed to convert input '%s' to int: %+v", inputString, err)
		}
		//
		// output the word
		//
		fmt.Printf("Output: %s\n\n", number(input))
	}
}
