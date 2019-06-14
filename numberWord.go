package main

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
	"teen",
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

func (num number) String() string {
	if num <= 999 {
		return under999(num)
	} else {
		word := ""
		position := 0
		for num > 0 {
			if num%1000 != 0 {
				s2 := under999(num % 1000)
				if position > 0 {
					s2 = s2 + " " + classifiers[position]
				}
				if word == "" {
					word = s2
				} else {
					word = s2 + " " + word
				}
			}
			num /= 1000
			position++
		}
		return word
	}
}

func under99(num number) string {
	if num < 20 {
		return singleAndTeenWords[num]
	} else {
		primaryPart := tens[num/10-2]
		secondPart := ""
		remainder := num % 10
		if remainder > 0 {
			secondPart = "-" + singleAndTeenWords[remainder]
		}
		return primaryPart + secondPart
	}
}

func under999(num number) string {
	if num <= 99 {
		return under99(num)
	} else {
		s1 := singleAndTeenWords[num/100] + " hundred"
		s2 := under99(num % 100)
		if num <= 99 {
			return s2
		}
		if num%100 == 0 {
			return s1
		} else {
			return s1 + " " + s2
		}
	}
}
