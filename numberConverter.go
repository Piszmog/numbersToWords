package main

import "strings"

const (
	space = " "
	dash  = "-"
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
	// using strings.Builder will help save on object creation and performance
	//
	wordBuilder := strings.Builder{}
	//
	// if under 1,000 handle it differently than if it was greater -- greater than 1,000 is more work
	//
	if num <= 999 {
		underOneThousandBuilder(num, &wordBuilder)
	} else {
		return overOneThousandBuilder(num, &wordBuilder)
	}
	return wordBuilder.String()
}

func underOneThousandBuilder(num number, wordBuilder *strings.Builder) {
	if num <= 99 {
		underOneHundredBuilder(num, wordBuilder)
	} else {
		hundredBuilder(num, wordBuilder)
		wordBuilder.WriteString(space)
		underOneHundredBuilder(num%100, wordBuilder)
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

func hundredBuilder(num number, wordBuilder *strings.Builder) {
	//
	// the special case when the number is in the hundreds
	//
	wordBuilder.WriteString(singleAndTeenWords[num/100])
	wordBuilder.WriteString(space)
	wordBuilder.WriteString(classifiers[0])
}

func overOneThousandBuilder(num number, wordBuilder *strings.Builder) string {
	word := ""
	position := 0
	for num > 0 {
		if num%1000 != 0 {
			underOneThousandBuilder(num%1000, wordBuilder)
			s2 := wordBuilder.String()
			wordBuilder.Reset()
			if position > 0 {
				s2 = s2 + space + classifiers[position]
			}
			if len(word) == 0 {
				word = s2
			} else {
				word = s2 + space + word
			}
		}
		num /= 1000
		position++
	}
	return word
}
