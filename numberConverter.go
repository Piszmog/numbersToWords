package main

import (
	"strings"
)

const (
	space    = " "
	dash     = "-"
	negative = "negative"
	zero     = "Zero"
)

var singleAndTeenWords = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tens = []string{
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

var classifiers = []string{
	"hundred",
	"thousand",
	"million",
	"billion",
	"trillion",
	"quadrillion",
	"quintillion",
}

type number int64

// String converts the number into the word equivalent. By implementing String(), the number type now
// implements fmt.Stringer. This means a special function does not need to be called to print the word.
func (num number) String() string {
	//
	// handle special case of 'zero'
	//
	if num == 0 {
		return zero
	}
	//
	// to save on object creation and performance, use strings.Builder
	//
	wordBuilder := strings.Builder{}
	//
	// handle negatives
	//
	numCopy := num
	if num < 0 {
		wordBuilder.WriteString(negative)
		numCopy = -num
	}
	//
	//
	// break the input into sections
	//
	var numberSections []number
	for numCopy > 0 {
		numberSections = append(numberSections, numCopy%1000)
		numCopy = numCopy / 1000
	}
	//
	// loop over the sections, building the word
	//
	for i := len(numberSections) - 1; i >= 0; i-- {
		numberSection := numberSections[i]
		//
		// If a section is 0, then skip it as there are no words for it
		//
		if numberSection == 0 {
			continue
		}
		//
		// If words have been added to the buffer, then add a space for the next section
		//
		if wordBuilder.Len() > 0 {
			wordBuilder.WriteString(space)
		}
		//
		// Convert the word to hundred based
		//
		convertHundred(numberSection, &wordBuilder)
		//
		// if the section is within the hundred section, then add the classifier
		//
		if i > 0 {
			wordBuilder.WriteString(space)
			wordBuilder.WriteString(classifiers[i])
		}
	}
	word := wordBuilder.String()
	return strings.ToUpper(word[:1]) + word[1:]
}

func convertHundred(num number, wordBuilder *strings.Builder) {
	if num <= 99 {
		underOneHundredBuilder(num, wordBuilder)
	} else {
		wordBuilder.WriteString(singleAndTeenWords[num/100])
		wordBuilder.WriteString(space)
		wordBuilder.WriteString(classifiers[0])
		//
		// If there is remainder, then get the next words
		//
		if num%100 != 0 {
			wordBuilder.WriteString(space)
			underOneHundredBuilder(num%100, wordBuilder)
		}
	}
}

func underOneHundredBuilder(num number, wordBuilder *strings.Builder) {
	if num < 20 {
		//
		// Since all numbers under 20 in english are special, just retrieve from the prebuilt array
		//
		wordBuilder.WriteString(singleAndTeenWords[num])
	} else {
		//
		// get the appropriate tens category
		//
		wordBuilder.WriteString(tens[num/10-2])
		remainder := num % 10
		//
		// if any remainder, then there are more words to convert
		//
		if remainder > 0 {
			wordBuilder.WriteString(dash + singleAndTeenWords[remainder])
		}
	}
}
