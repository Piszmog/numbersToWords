# Numbers to Words
[![Build Status](https://travis-ci.org/Piszmog/numbersToWords.svg?branch=develop)](https://travis-ci.org/Piszmog/gotoggle)
[![Coverage Status](https://coveralls.io/repos/github/Piszmog/numbersToWords/badge.svg?branch=develop)](https://coveralls.io/github/Piszmog/numbersToWords?branch=develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/Piszmog/numbersToWords)](https://goreportcard.com/report/github.com/Piszmog/numbersToWords)
[![GitHub release](https://img.shields.io/github/release/Piszmog/numbersToWords.svg)](https://github.com/Piszmog/numbersToWords/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Application for converting numbers into their word equivalent - i.e  `1` -> `One`

## Running
Download the appropriate executable for the Operating System that it will be ran on (Linux, Mac, Windows). Then simply run 
the executable with the `-f` flag to specify the input file to read.

```bash
$ ./numberToWord-mac -f {file to read}
```

###### Example
```bash
$ ./numberToWord-mac -f ~/Documents/input.txt
```

## File
The application reads the provided text file (`.txt`) and converts the number on each line to their word equivalent.

### Format
The expected format of the file to read is each line is the number to convert

###### Example Input
```text
1
10
4695
759672
```

###### Example Output
```text
Input: 1
Output: One
Input: 10
Output: Ten
Input: 4695
Output: Four thousand six hundred ninety-five
Input: 759672
Output: Seven hundred fifty-nine thousand six hundred seventy-two
```

### Invalid Inputs
If a line in the file contains `,`, `.` or letters, an error will be printed out - `Error: Input '1.456' is not a valid input.`
