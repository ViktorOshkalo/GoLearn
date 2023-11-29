package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Hello!")

	text, err := os.ReadFile("numbers.txt")
	if err != nil {
		panic("unable to read file")
	}

	fmt.Println("Input text:")
	fmt.Println(string(text))

	rexp := regexp.MustCompile(`(\+?(?P<code_country>\d{2})[. -]?)?\(?(?P<code_operator>\d{3})\)?[. -]?(?P<number_p1>\d{3})[. -]?(?P<number_p2>\d{2})[. -]?(?P<number_p3>\d{2})`)

	res := rexp.FindAllString(string(text), -1)
	if res == nil {
		fmt.Println("\nNOT found.")
	}

	fmt.Println("\nNumbers found:")
	for _, r := range res {
		fmt.Println(string(r))
	}

	matches := rexp.FindAllStringSubmatch(string(text), -1)
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

	fmt.Println("Exit")
}
