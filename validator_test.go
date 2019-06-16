package main

import (
	"testing"
)

func TestIsStringValid(t *testing.T) {
	if !numberValidator.isStringValid("1") {
		t.Fatal("Expected string '1' to be valid")
	}
	if !numberValidator.isStringValid("-1") {
		t.Fatal("Expected string '-1' to be valid")
	}
	if numberValidator.isStringValid("1.345") {
		t.Fatal("Expected string '1.345' to be not valid")
	}
	if numberValidator.isStringValid("1.add") {
		t.Fatal("Expected string '1.add' to be not valid")
	}
	if numberValidator.isStringValid("person") {
		t.Fatal("Expected string '1.add' to be not valid")
	}
	if numberValidator.isStringValid("--1") {
		t.Fatal("Expected string '--1' to be not valid")
	}
	if numberValidator.isStringValid("1-") {
		t.Fatal("Expected string '1-' to be not valid")
	}
	if numberValidator.isStringValid("1,000") {
		t.Fatal("Expected string '1,000' to be not valid")
	}
}
