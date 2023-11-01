package main

import (
	"bufio"
	"fmt"
	"os"

	CsvGenerator "iprice-assessment-test/csv_generator"
	StringManipulation "iprice-assessment-test/string_manipulation"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Input a string : ")
	scanner.Scan()

	inputtedString := scanner.Text()
	uppercase := StringManipulation.GetUppercase(inputtedString)
	alternatecase := StringManipulation.GetAlternateCase(inputtedString)
	fileCsv := CsvGenerator.GenerateCSV(inputtedString)

	fmt.Println("Uppercase", uppercase)
	fmt.Println("Alternatecase", alternatecase)
	fmt.Println("CSV generated", fileCsv)
}
