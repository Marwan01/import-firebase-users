package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open the file
	csvfile, err := os.Open("save_file.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	var emails []string
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		emails = append(emails, record[1])
	}

	f, err := os.Create("emails.txt")

	if err != nil {
		log.Fatal(err)
	}
	_, _ = f.WriteString("Email Address\n")

	defer f.Close()
	for i := 0; i < len(emails); i++ {
		_, err2 := f.WriteString(emails[i])
		_, _ = f.WriteString("\n")
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	fmt.Println("Successfully exported all firebase users to emails.txt")
}
