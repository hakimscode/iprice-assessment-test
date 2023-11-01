package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Input a string : ")
	scanner.Scan()

	inputtedString := scanner.Text()
	fmt.Printf("%s\n", inputtedString)
}
