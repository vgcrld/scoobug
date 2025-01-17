package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	var file *os.File
	var ok error

	filename := flag.String("file", "data.csv", "CSV file to read")
	flag.Parse()

	if file, ok = os.Open(*filename); ok != nil {
		fmt.Println("Error opening file:", ok)
		return
	}

	defer func() { file.Close(); fmt.Println("File closed") }()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	uniqueValues := make(map[string]bool)
	for _, record := range records[1:] { // Skip the header row
		uniqueValues[record[0]] = true
	}

	for value := range uniqueValues {
		fmt.Println(value)
	}
}
