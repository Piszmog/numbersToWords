package main

import "regexp"

var numberValidator = validator{regex: regexp.MustCompile("^-?[0-9]+$")}

type validator struct {
	regex *regexp.Regexp
}

func (validator validator) isStringValid(num string) bool {
	return validator.regex.MatchString(num)
}
