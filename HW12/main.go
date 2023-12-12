package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// strategy
type ITextProcessingStrategy interface {
	SetParams()
	Process(text string) string
}

type TextReplacerStrategy struct {
	From string
	To   string
}

func (r TextReplacerStrategy) Process(text string) string {
	return strings.ReplaceAll(text, r.From, r.To)
}

func (r *TextReplacerStrategy) SetParams() {
	var oldWord, newWord string
	fmt.Println("Enter old word: ")
	fmt.Scanln(&oldWord)
	fmt.Println("Enter new word: ")
	fmt.Scanln(&newWord)
	r.From = oldWord
	r.To = newWord
}

type SpaceCleanerStrategy struct {
	SelectorRegex string
	To            string
}

func (r SpaceCleanerStrategy) Process(text string) string {
	re := regexp.MustCompile(r.SelectorRegex)
	return re.ReplaceAllString(text, r.To)
}

func (r *SpaceCleanerStrategy) SetParams() {
	r.SelectorRegex = `\s+`
	r.To = " "
}

// text processor
type ITextProcessor interface {
	ProcessText(text string) string
}

// some ITextProcessor implementation
type TextProcessor struct {
}

func (p TextProcessor) ProcessText(text string) string {
	return text
}

// decorator
type TextProcessorDecorator struct {
	base     ITextProcessor
	strategy ITextProcessingStrategy
}

func (p TextProcessorDecorator) ProcessText(text string) string {
	if p.base != nil {
		text = p.base.ProcessText(text)
	}
	return p.strategy.Process(text)
}

func GetNextTextProcessor(command string, baseProcessor ITextProcessor) (ITextProcessor, error) {
	switch command {
	case "space":
		strategy := SpaceCleanerStrategy{}
		strategy.SetParams()
		return TextProcessorDecorator{base: baseProcessor, strategy: &strategy}, nil
	case "word":
		strategy := TextReplacerStrategy{}
		strategy.SetParams()
		return TextProcessorDecorator{base: baseProcessor, strategy: &strategy}, nil
	default:
		return nil, fmt.Errorf("unknown command: %s", command)
	}
}

var commands map[string]string = map[string]string{
	"space": "add space cleaner",
	"word":  "add word replacer",
	"done":  "start text modifying",
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

	var processor ITextProcessor = TextProcessor{}
	for {
		fmt.Println("\nEnter command: ")
		var command string
		fmt.Scanln(&command)

		if command == "done" {
			text = processor.ProcessText(text)
			fmt.Println("Modified text:")
			fmt.Println(text)
			return
		}

		var err error
		processor, err = GetNextTextProcessor(command, processor)
		if err != nil {
			fmt.Println(err)
			continue
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
