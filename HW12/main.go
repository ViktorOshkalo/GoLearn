package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type ITextModifier interface {
	ModifyText(text string) string
}

// base
type BaseDecorator struct {
	TextModifier ITextModifier
}

func (b BaseDecorator) ModifyText(text string) string {
	if b.TextModifier != nil {
		text = b.TextModifier.ModifyText(text)
	}
	return text
}

// space cleaner
type SpaceCleanerDecorator struct {
	BaseDecorator
}

func (w SpaceCleanerDecorator) ModifyText(text string) string {
	text = w.BaseDecorator.ModifyText(text)
	re := regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")
	return text
}

// word replacer
type WordReplacerDecorator struct {
	BaseDecorator
	OldWord string
	NewWord string
}

func (w WordReplacerDecorator) ModifyText(text string) string {
	text = w.BaseDecorator.ModifyText(text)
	text = strings.ReplaceAll(text, w.OldWord, w.NewWord)
	return text
}

var commands map[string]string = map[string]string{
	"space": "add space cleaner",
	"word":  "add word replacer",
	"done":  "start text modifiing",
}

func printCommands() {
	fmt.Println("Available commands:")
	for key, value := range commands {
		fmt.Printf("%s - %s\n", key, value)
	}
}

func main() {
	fmt.Println("Hello")

	var text = ReadFile("file.txt")
	fmt.Println("Original Text:")
	fmt.Println(text)

	fmt.Printf("\nEnter a commands to build text modifier.\n")
	printCommands()

	var worker ITextModifier = BaseDecorator{TextModifier: nil}
	for {
		fmt.Println("\nEnter command: ")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "space":
			worker = SpaceCleanerDecorator{BaseDecorator: BaseDecorator{TextModifier: worker}}
			fmt.Println("Space cleaner added")
		case "word":
			var oldWord, newWord string
			fmt.Println("Enter old word: ")
			fmt.Scanln(&oldWord)
			fmt.Println("Enter new word: ")
			fmt.Scanln(&newWord)
			worker = WordReplacerDecorator{BaseDecorator: BaseDecorator{TextModifier: worker}, OldWord: oldWord, NewWord: newWord}
			fmt.Println("Word replacer added")
		case "done":
			text = worker.ModifyText(text)
			fmt.Println("Modified text:")
			fmt.Println(text)
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}

func ReadFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(content)
}
