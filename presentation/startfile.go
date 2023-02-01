package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//file, err := os.Open("model/file.csv") // opens csv file
	file, err := os.Open(commandLineArg()[1]) // opens csv file
	reader := readFile(file, err)             // reads it
	ourTable, _ := reader.ReadAll()

	for y := 1; y < len(ourTable); y++ {
		for x := 1; x < len(ourTable[0]); x++ {
			cur := ourTable[y][x]
			if strings.Contains(string(cur[0]), "=") {
				ourTable[y][x] = getResult(&ourTable[y][x], ourTable)
			}
		}
	}

	//Printing in console whole table
	for i := 0; i < len(ourTable); i++ {
		fmt.Println(strings.Join(ourTable[i], ","))
	}
}

// Reading file from csv file and returning
// pointer to  the *cvs.Reader
func readFile(file *os.File, err error) *csv.Reader {

	if err != nil {
		panic(err)
	}

	cvsFile := csv.NewReader(file)

	return cvsFile
}

// Looking for specific position in array where contains our string
func findCell(xAddr, yAddr string, ourTable [][]string) *string {

	for y := range ourTable {
		if ourTable[y][0] != yAddr {
			continue
		}
		for x := range ourTable {
			if ourTable[0][x] == xAddr {
				return &ourTable[y][x]
			}
		}
	}

	return nil
}

func commandLineArg() []string {

	//argsWithProg := os.Args
	filePath := os.Args

	return filePath
}

// Splitting string (formula, example "=A1+C2")
// into array of 2 substrings [A1] and [C2].
// Up next, we are working on each string.
// We separate symbol and numbers (example: [A1] -> [A], [1]
// Searching them in our columns and rows (but specifically
// in first column and first row).
// Returning the calculated formula.
func getResult(str *string, ourTable [][]string) string {
	cur := *str
	var op string            //operation
	var twoStrings [2]string //num1, num2,

	if strings.Contains(cur, "=") {

		for opNum := 0; opNum < len(*str); opNum++ {

			if cur[opNum] == '+' || cur[opNum] == '-' || cur[opNum] == '*' || cur[opNum] == '/' {

				twoStrings[0] = cur[1:opNum]
				twoStrings[1] = cur[opNum+1:]
				op = string(cur[opNum])
				break
			}

		}

		for i := 0; i < 2; i++ {
			if !isNumber(twoStrings[i]) {
				var x, y string
				j := 0

				for j < len(twoStrings[i]) {

					if isNumber(string(twoStrings[i][j])) {
						break
					}
					j++
				}
				x = twoStrings[i][0:j]
				y = twoStrings[i][j:len(twoStrings[i])]

				var cell *string = findCell(x, y, ourTable)
				var value string = *cell
				if strings.Contains(value, "=") {
					*cell = getResult(cell, ourTable)
				}
				twoStrings[i] = *cell
			}

		}

	}

	return calc(twoStrings[0], twoStrings[1], op)
}

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

// Checking if string contains number
func isNumber(str string) bool {

	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true

}
