package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var text []string = []string{
	"How do you feel?",
	"I'm feel",
	"I'm very feel",
}

func searchUserText(searchString string) (searchResult []string) {
	fmt.Println("searching...")
	result := []string{}
	for _, row := range text {
		if row == searchString {
			result = append(result, row)
		} else if strings.Contains(row, searchString) {
			result = append(result, row)
		}
	}
	return result
}

func printText(text []string) {
	for _, row := range text {
		fmt.Println(row)
	}
}

func main() {
	fmt.Println("User text:")
	printText(text)

	var searchString string
	fmt.Println("\nEnter a search string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		searchString = scanner.Text()
	}

	searchResult := searchUserText(searchString)
	fmt.Println("\nSearch result: ")
	printText(searchResult)
}
