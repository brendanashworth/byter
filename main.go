package main

import (
	table "github.com/crackcomm/go-clitable"
	"io/ioutil"
	"strconv"
	"fmt"
	"os"
)

func main() {
	// ensure provided two arguments (./byter <file>)
	if len(os.Args) != 2 {
		fmt.Printf("\n  Usage: %s <binary file>\n\n", os.Args[0]) // os.Args[0] is location of binary
		return
	}

	// read file, os.Args[1] is binary file
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error occurred while reading file: %s\n", err.Error())
		return
	}

	// count occurrences then strip by threshold of 1%
	bytes, occurrenceMap := CountOccurrences(data)
	totalBytes := float64(bytes)

	occurrenceMap = StripOccurrences(occurrenceMap, totalBytes, 0.01)

	// convert to map[string]int for cli table
	tableData := make(map[string]interface{})

	for key, value := range occurrenceMap {
		tableData[fmt.Sprintf("0x%X", key)] = strconv.Itoa(int((float64(value) / totalBytes) * 100)) + "%" // gives nice value percentage
	}

	// create table with remaining occurrences
	table.PrintHorizontal(tableData)

	fmt.Printf("\nParsed %d bytes.\n", bytes)
}
