package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("model/file.csv") // opens csv file
	reader := readFile(file, err)          // reads it
	ourTable, _ := reader.ReadAll()

	for y := 1; y < len(ourTable); y++ {
		for x := 1; x < len(ourTable[0]); x++ {
			cur := ourTable[y][x]
			if strings.Contains(string(cur[0]), "=") {
				ourTable[y][x] = getResult(&ourTable[y][x], ourTable)
				println()
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
