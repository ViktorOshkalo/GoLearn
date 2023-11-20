package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MinMax struct {
	Min int
	Max int
}

func generateNumbersAndPrintMinMax(numbersCount int, maxNumber int, channelNumbers chan int, channelMinMax chan MinMax, channelDone chan bool) {
	for i := 0; i < numbersCount; i++ {
		num := rand.Intn(maxNumber)
		fmt.Println("Sending: ", num)
		time.Sleep(time.Millisecond)
		channelNumbers <- num
	}
	close(channelNumbers)

	minMax := <-channelMinMax

	fmt.Println("Min: ", minMax.Min)
	fmt.Println("Max: ", minMax.Max)

	channelDone <- true
}

func findMinMax(channelNumbers chan int, channelMinMax chan MinMax) {
	var numbers []int
	for n := range channelNumbers {
		fmt.Println("Receiving: ", n)
		numbers = append(numbers, n)
	}

	if len(numbers) == 0 {
		panic("No numbers recieved!")
	}

	fmt.Println("Numbers received: ", numbers)

	var min, max int = numbers[0], numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
	}

	channelMinMax <- MinMax{Min: min, Max: max}

	close(channelMinMax)
}

func main() {
	fmt.Println("Go!")

	numbersCount := 0
	maxValue := 100

	channelNumbers := make(chan int, numbersCount)
	channelMinMax := make(chan MinMax)
	channelDone := make(chan bool)

	go generateNumbersAndPrintMinMax(numbersCount, maxValue, channelNumbers, channelMinMax, channelDone)
	go findMinMax(channelNumbers, channelMinMax)

	<-channelDone
}
