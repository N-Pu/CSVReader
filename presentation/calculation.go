package main

import (
	"strconv"
)

// Calculation and adding one of 4 operations between values.
func calc(str1, str2, op string) string {

	var result string

	switch op {
	case "+":
		res1, _ := strconv.Atoi(str1)
		res2, _ := strconv.Atoi(str2)
		result = strconv.Itoa(res1 + res2)
		break
	case "-":
		res1, _ := strconv.Atoi(str1)
		res2, _ := strconv.Atoi(str2)
		result = strconv.Itoa(res1 - res2)
		break
	case "*":
		res1, _ := strconv.Atoi(str1)
		res2, _ := strconv.Atoi(str2)
		result = strconv.Itoa(res1 * res2)
		break
	case "/":
		res1, _ := strconv.Atoi(str1)
		res2, _ := strconv.Atoi(str2)
		result = strconv.Itoa(res1 / res2)
		break
	}

	return result

}
