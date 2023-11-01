package csv_generator

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	StringManipulation "iprice-assessment-test/string_manipulation"
)

func GenerateCSV(str string) string {
	uppercase := StringManipulation.GetUppercase(str)
	alternatecase := StringManipulation.GetAlternateCase(str)

	arrUppercase := strings.Split(uppercase, "")
	arrAlternatecase := strings.Split(alternatecase, "")

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	filename := path + "/" + str + ".csv"

	csvFile, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed creating csv: %v", err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	var data [][]string
	data = append(data, arrUppercase)
	data = append(data, arrAlternatecase)

	w.WriteAll(data)

	return filename
}
