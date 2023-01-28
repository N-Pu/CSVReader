package main

import "unicode"

// Checking if string contains number
func isNumber(str string) bool {

	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true

}
