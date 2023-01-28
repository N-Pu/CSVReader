package main

import "strings"

//Splitting string (formula, example "=A1+C20"
//into array of 2 strings [A1] and [C2].
//Up next, we are working on each string.
//We separate symbol and numbers (example: [A1] -> [A], [1]
//Searching them in our columns and rows (but specifically
//in first column and first row).
//Returning the calculated formula.

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
