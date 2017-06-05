package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)



func csvUtility() {
	// reading data from JSON File
	// Unmarshal JSON data

	// Create a csv file
	f, err := os.Create("./contact.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// Write Unmarshaled json data to CSV file
	w := csv.NewWriter(f)
	for _, obj := range contacts {
		var record []string
		record = append(record, strconv.Itoa(obj.Id))
		record = append(record, obj.FirstName)
		record = append(record, obj.LastName)
		record = append(record, obj.Email)
		record = append(record, obj.PhoneNumber)
		w.Write(record)
	}
	w.Flush()
}