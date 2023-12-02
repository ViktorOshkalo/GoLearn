package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Hello!")

	reader, err := ReadFileByLineBatches("numbers.txt", 3)
	if err != nil {
		panic(err)
	}

	rexp := regexp.MustCompile(`(\+?(?P<code_country>\d{2})[. -]?)?\(?(?P<code_operator>\d{3})\)?[. -]?(?P<number_p1>\d{3})[. -]?(?P<number_p2>\d{2})[. -]?(?P<number_p3>\d{2})`)

	for text := range reader {

		fmt.Println("\nText batch received: ")
		fmt.Println(text)

		matches := rexp.FindAllStringSubmatch(string(text), -1)

		fmt.Println("\nAll findings:")
		for _, match := range matches {
			fmt.Println(match[0])
		}

		var numberInfos []map[string]string
		for _, match := range matches {
			subMatchMap := make(map[string]string)
			for i, name := range rexp.SubexpNames() {
				if i != 0 && name != "" {
					subMatchMap[name] = string(match[i])
				}
			}
			numberInfos = append(numberInfos, subMatchMap)
		}

		fmt.Println("\nNumbers normalized:")
		for _, numberInfo := range numberInfos {
			if numberInfo["code_country"] != "" {
				fmt.Printf("+%s (%s) %s-%s\n",
					numberInfo["code_country"],
					numberInfo["code_operator"],
					numberInfo["number_p1"],
					numberInfo["number_p2"]+numberInfo["number_p3"])
			} else {
				fmt.Printf("(%s) %s-%s\n",
					numberInfo["code_operator"],
					numberInfo["number_p1"],
					numberInfo["number_p2"]+numberInfo["number_p3"])
			}
		}
	}
	fmt.Println("Exit")
}

// lets say file is super big
func ReadFileByLineBatches(fileName string, linesBatchSize int) (<-chan string, error) {
	out := make(chan string)
	file, err := os.Open(fileName)
	if err != nil {
		return out, err
	}

	go func() {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			linesOut := scanner.Text()
			for i := 1; i < linesBatchSize && scanner.Scan(); i++ {
				linesOut += "\n" + scanner.Text()
			}
			out <- linesOut
		}
		close(out)
	}()

	return out, nil
}
