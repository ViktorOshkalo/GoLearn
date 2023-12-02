package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Hello!")

	text, err := os.ReadFile("text.txt")
	if err != nil {
		panic("unable to read file")
	}

	fmt.Println("Input text:")
	fmt.Println(string(text))

	rexp := regexp.MustCompile(`(^|[^\p{L}])([аоуіие]+\p{L}+[бвгґджзйклмнпрстфхцчшщ]+)([^\p{L}]|$)`)
	matches := rexp.FindAllStringSubmatch(string(text), -1)

	fmt.Println("\nAll found words - begins vowel, finishes consonant:")
	for _, match := range matches {
		fmt.Println(match[2])
	}

	rexp2 := regexp.MustCompile(`(^|[^\p{L}])([бвгґджзйклмнпрстфхцчшщ]+\p{L}+[аоуіие]+)([^\p{L}]|$)`)
	matches2 := rexp2.FindAllStringSubmatch(string(text), -1)

	fmt.Println("\nAll found words - begins consonant, finishes vowel:")
	for _, match := range matches2 {
		fmt.Println(match[2])
	}
}
