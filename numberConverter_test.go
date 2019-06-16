package main

import (
	"testing"
)

func TestNumber_String(t *testing.T) {
	word := number(-10).String()
	if word != "Negative ten" {
		t.Fatalf("expected the word 'Negative ten' but was converted to '%s'", word)
	}
	word = number(0).String()
	if word != "Zero" {
		t.Fatalf("expected the word 'Zero' but was converted to '%s'", word)
	}
	word = number(1).String()
	if word != "One" {
		t.Fatalf("expected the word 'One' but was converted to '%s'", word)
	}
	word = number(10).String()
	if word != "Ten" {
		t.Fatalf("expected the word 'Ten' but was converted to '%s'", word)
	}
	word = number(15).String()
	if word != "Fifteen" {
		t.Fatalf("expected the word 'Ten' but was converted to '%s'", word)
	}
	word = number(35).String()
	if word != "Thirty-five" {
		t.Fatalf("expected the word 'Thirty-five' but was converted to '%s'", word)
	}
	word = number(99).String()
	if word != "Ninety-nine" {
		t.Fatalf("expected the word 'Ninety-nine' but was converted to '%s'", word)
	}
	word = number(100).String()
	if word != "One hundred" {
		t.Fatalf("expected the word 'One hundred' but was converted to '%s'", word)
	}
	word = number(102).String()
	if word != "One hundred two" {
		t.Fatalf("expected the word 'One hundred two' but was converted to '%s'", word)
	}
	word = number(155).String()
	if word != "One hundred fifty-five" {
		t.Fatalf("expected the word 'One hundred fifty-five' but was converted to '%s'", word)
	}
	word = number(999).String()
	if word != "Nine hundred ninety-nine" {
		t.Fatalf("expected the word 'Nine hundred ninety-nine' but was converted to '%s'", word)
	}
	word = number(1000).String()
	if word != "One thousand" {
		t.Fatalf("expected the word 'One thousand' but was converted to '%s'", word)
	}
	word = number(1012).String()
	if word != "One thousand twelve" {
		t.Fatalf("expected the word 'One thousand twelve' but was converted to '%s'", word)
	}
	word = number(1708).String()
	if word != "One thousand seven hundred eight" {
		t.Fatalf("expected the word 'One thousand seven hundred eight' but was converted to '%s'", word)
	}
	word = number(8708).String()
	if word != "Eight thousand seven hundred eight" {
		t.Fatalf("expected the word 'Eight thousand seven hundred eight' but was converted to '%s'", word)
	}
	word = number(30708).String()
	if word != "Thirty thousand seven hundred eight" {
		t.Fatalf("expected the word 'Thirty thousand seven hundred eight' but was converted to '%s'", word)
	}
	word = number(55708).String()
	if word != "Fifty-five thousand seven hundred eight" {
		t.Fatalf("expected the word 'Fifty-five thousand seven hundred eight' but was converted to '%s'", word)
	}
	word = number(100001).String()
	if word != "One hundred thousand one" {
		t.Fatalf("expected the word 'One hundred thousand one' but was converted to '%s'", word)
	}
	word = number(1000001).String()
	if word != "One million one" {
		t.Fatalf("expected the word 'One million one' but was converted to '%s'", word)
	}
	word = number(1000000001).String()
	if word != "One billion one" {
		t.Fatalf("expected the word 'One billion one' but was converted to '%s'", word)
	}
	word = number(1000000000001).String()
	if word != "One trillion one" {
		t.Fatalf("expected the word 'One trillion one' but was converted to '%s'", word)
	}
	word = number(1000000000000001).String()
	if word != "One quadrillion one" {
		t.Fatalf("expected the word 'One quadrillion one' but was converted to '%s'", word)
	}
	word = number(1000000000000000001).String()
	if word != "One quintillion one" {
		t.Fatalf("expected the word 'One quintillion one' but was converted to '%s'", word)
	}
	//
	// Let's check the largest int64 number possible
	//
	word = number(9223372036854775807).String()
	if word != "Nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred seven" {
		t.Fatalf("expected the word 'Nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred seven' but was converted to '%s'", word)
	}
}
