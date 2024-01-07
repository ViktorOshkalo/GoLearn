package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// text processor interface
type ITextProcessor interface {
	Process(text string) map[string]int
}

// strtegies
type WordsCounter struct{}

func (wc WordsCounter) Process(text string) map[string]int {
	fmt.Printf("\nIncoming text:\n\n%s", text)
	words := strings.Fields(text)

	out := map[string]int{
		"words count": len(words),
	}
	return out
}

type WordsEntriesCounter struct{}

func (wc WordsEntriesCounter) Process(text string) map[string]int {
	fmt.Printf("\nIncoming text:\n\n%s", text)

	re := regexp.MustCompile(`\w+`)
	matchWords := re.FindAllString(text, -1)

	stats := make(map[string]int)
	for _, word := range matchWords {
		stats[word]++
	}

	return stats
}

// decorators
// space cleaner
type SpaceCleaner struct {
	base ITextProcessor
}

func (sr SpaceCleaner) Process(text string) map[string]int {
	re := regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")
	if sr.base != nil {
		return sr.base.Process(text)
	} else {
		panic("decorator base is not set")
	}
}

// word replacer
type WordReplacer struct {
	base ITextProcessor
	From string
	To   string
}

func (r WordReplacer) Process(text string) map[string]int {
	text = strings.ReplaceAll(text, r.From, r.To)
	if r.base != nil {
		return r.base.Process(text)
	} else {
		panic("decorator base is not set")
	}
}

// main
func main() {
	var text = ReadFile("file.txt")
	fmt.Println("Original Text:")
	fmt.Println(text)

	// test case
	fmt.Println("\n\nTest case mode")
	processor := SpaceCleaner{base: WordReplacer{From: "sun", To: "SUN", base: WordsEntriesCounter{}}}
	res := processor.Process(text)
	PrintResult(res)

	// user input
	fmt.Println("\n\nUser input mode")
	strategy := GetStrategy()
	userProcessor := GetDecorator(strategy)
	res = userProcessor.Process(text)
	PrintResult(res)
}

// helper methods
func GetStrategy() ITextProcessor {
	fmt.Println("Strategy setup")
	fmt.Println("Available strategies:")
	fmt.Println("word: Calculate words count")
	fmt.Println("entry: Calculate words entries count")

	for {
		var command string
		fmt.Print("Enter a command: ")
		fmt.Scanln(&command)

		switch command {
		case "word":
			return WordsCounter{}
		case "entry":
			return WordsEntriesCounter{}
		default:
			fmt.Println("Strategy not exists. Enter correct command")
			continue
		}
	}
}

func GetDecorator(base ITextProcessor) ITextProcessor {
	fmt.Println("Decorators setup")
	fmt.Println("Available decorators:")
	fmt.Println("space: Space remover")
	fmt.Println("word: Word replacer")
	fmt.Println("done: Exit decorators setup")

	var processor ITextProcessor = base
	for {
		var command string
		fmt.Print("Enter a command: ")
		fmt.Scanln(&command)

		switch command {
		case "space":
			processor = SpaceCleaner{base: processor}
		case "word":
			var oldWord, newWord string
			fmt.Println("Enter old word: ")
			fmt.Scanln(&oldWord)
			fmt.Println("Enter new word: ")
			fmt.Scanln(&newWord)
			processor = WordReplacer{base: processor, From: oldWord, To: newWord}
		case "done":
			return processor
		default:
			fmt.Println("Wrong command")
			continue
		}
	}
}

func PrintResult(res map[string]int) {
	fmt.Println("\n\nResult:")
	for k, v := range res {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func ReadFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(content)
}
