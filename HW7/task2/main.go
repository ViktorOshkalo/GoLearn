package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbersAndPrintMinMax(numbersCount int, maxNumber int, channelNumbers chan int, channelMinMax chan int) {
	for i := 0; i < numbersCount; i++ {
		num := rand.Intn(maxNumber)
		fmt.Println("Sending: ", num)
		time.Sleep(time.Millisecond)
		channelNumbers <- num
	}
	close(channelNumbers)

	min := <-channelMinMax
	max := <-channelMinMax

	if min > max {
		panic("wrong min max values")
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

func findMinMax(channelNumbers chan int, channelMinMax chan int) {
	var numbers []int
	for n := range channelNumbers {
		fmt.Println("Receiving: ", n)
		numbers = append(numbers, n)
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

	channelMinMax <- min
	channelMinMax <- max

	close(channelMinMax)
}

func main() {
	fmt.Println("Go!")

	numbersCount := 10
	maxValue := 100

	channelNumbers := make(chan int, numbersCount)
	channelMinMax := make(chan int)

	go generateNumbersAndPrintMinMax(numbersCount, maxValue, channelNumbers, channelMinMax)
	go findMinMax(channelNumbers, channelMinMax)

	time.Sleep(time.Second)
}
