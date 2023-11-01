package csv_generator

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"testing"

	StringManipulation "iprice-assessment-test/string_manipulation"
)

type Test struct {
	name  string
	input string
}

func TestGenerateCSV(t *testing.T) {

	tests := []Test{
		{
			name:  "Test generate csv with normal string",
			input: "herihakim",
		},
		{
			name:  "Test generate csv input with space",
			input: "heri hakim",
		},
		{
			name:  "Test generate csv input containing uppercase",
			input: "Heri Hakim",
		},
	}

	for _, test := range tests {
		generatedCsv := GenerateCSV(test.input)

		fileCsv, err := os.Open(generatedCsv)
		if err != nil {
			t.Errorf("CSV not found")
			continue
		}
		defer fileCsv.Close()

		csvReader := csv.NewReader(fileCsv)
		data, err := csvReader.ReadAll()
		if err != nil {
			t.Errorf("Error read csv")
		}

		expectedUppercase := StringManipulation.GetUppercase(test.input)
		expectedAlternatecase := StringManipulation.GetAlternateCase(test.input)

		if len(data) == 0 {
			t.Error("Failed, csv has no data")
		}

		for i, item := range data {
			result := strings.Join(item, "")
			if i == 0 {
				if result != expectedUppercase {
					t.Errorf("Failed, expected output is %s but the result is %s", expectedUppercase, result)
				}
			}

			if i == 1 {
				if result != expectedAlternatecase {
					t.Errorf("Failed, expected output is %s but the result is %s", expectedAlternatecase, result)
				}
			}

		}

		if err := os.Remove(generatedCsv); err != nil {
			log.Printf("Failed to delete: %v", err)
		}
	}
}
