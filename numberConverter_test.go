package main

import (
	"testing"
)

func TestNumber_String(t *testing.T) {
	word := number(1).String()
	if word != "one" {
		t.Fatalf("expected the word 'one' but was converted to '%s'", word)
	}
	word = number(10).String()
	if word != "ten" {
		t.Fatalf("expected the word 'ten' but was converted to '%s'", word)
	}
	word = number(15).String()
	if word != "fifteen" {
		t.Fatalf("expected the word 'ten' but was converted to '%s'", word)
	}
	word = number(35).String()
	if word != "thirty-five" {
		t.Fatalf("expected the word 'thirty-five' but was converted to '%s'", word)
	}
	word = number(99).String()
	if word != "ninety-nine" {
		t.Fatalf("expected the word 'ninety-nine' but was converted to '%s'", word)
	}
	word = number(102).String()
	if word != "one hundred two" {
		t.Fatalf("expected the word 'one hundred two' but was converted to '%s'", word)
	}
	word = number(100).String()
	if word != "one hundred" {
		t.Fatalf("expected the word 'one hundred' but was converted to '%s'", word)
	}
	word = number(155).String()
	if word != "one hundred fifty-five" {
		t.Fatalf("expected the word 'one hundred fifty-five' but was converted to '%s'", word)
	}
	word = number(999).String()
	if word != "nine hundred ninety-nine" {
		t.Fatalf("expected the word 'nine hundred ninety-nine' but was converted to '%s'", word)
	}
	word = number(1000).String()
	if word != "one thousand" {
		t.Fatalf("expected the word 'one thousand' but was converted to '%s'", word)
	}
	word = number(1012).String()
	if word != "one thousand twelve" {
		t.Fatalf("expected the word 'one thousand twelve' but was converted to '%s'", word)
	}
	word = number(1708).String()
	if word != "one thousand seven hundred eight" {
		t.Fatalf("expected the word 'one thousand seven hundred eight' but was converted to '%s'", word)
	}
	word = number(8708).String()
	if word != "eight thousand seven hundred eight" {
		t.Fatalf("expected the word 'eight thousand seven hundred eight' but was converted to '%s'", word)
	}
	word = number(30708).String()
	if word != "thirty thousand seven hundred eight" {
		t.Fatalf("expected the word 'thirty thousand seven hundred eight' but was converted to '%s'", word)
	}
	word = number(55708).String()
	if word != "fifty-five thousand seven hundred eight" {
		t.Fatalf("expected the word 'fifty-five thousand seven hundred eight' but was converted to '%s'", word)
	}
	word = number(100001).String()
	if word != "one hundred thousand one" {
		t.Fatalf("expected the word 'one hundred thousand one' but was converted to '%s'", word)
	}
	word = number(1000001).String()
	if word != "one million one" {
		t.Fatalf("expected the word 'one million one' but was converted to '%s'", word)
	}
	word = number(1000000001).String()
	if word != "one billion one" {
		t.Fatalf("expected the word 'one billion one' but was converted to '%s'", word)
	}
	word = number(1000000000001).String()
	if word != "one trillion one" {
		t.Fatalf("expected the word 'one trillion one' but was converted to '%s'", word)
	}
	word = number(1000000000000001).String()
	if word != "one quadrillion one" {
		t.Fatalf("expected the word 'one quadrillion one' but was converted to '%s'", word)
	}
	word = number(1000000000000000001).String()
	if word != "one quintillion one" {
		t.Fatalf("expected the word 'one quintillion one' but was converted to '%s'", word)
	}
	//
	// Let's check the largest int64 number possible
	//
	word = number(9223372036854775807).String()
	if word != "nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred seven" {
		t.Fatalf("expected the word 'nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred seven' but was converted to '%s'", word)
	}
}
