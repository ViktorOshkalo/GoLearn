package main

import (
	"fmt"
	"math/rand"
)

func printAverage(channelAvg chan float32, channelDone chan bool) {
	fmt.Println("printAvg started")
	avg := <-channelAvg
	fmt.Println("Average value is: ", avg)
	channelDone <- true
}

func calcAverage(channelNumbers chan int, channelAvg chan float32) {
	fmt.Println("calcAverage started")
	var sum int
	var counter int
	for val := range channelNumbers {
		fmt.Println("recieving from numbers channel: ", val)
		sum += val
		counter++
	}

	avg := float32(sum) / float32(counter)
	fmt.Println("sending average value into channel: ", avg)
	channelAvg <- avg
	close(channelAvg)
}

func generateNumbers(numbersCount int, maxNumber int, channelNumbers chan int) {
	fmt.Println("generateNumbers started")
	for i := 0; i < numbersCount; i++ {
		num := rand.Intn(maxNumber)
		fmt.Println("sending into numbers channel: ", num)
		channelNumbers <- num
	}
	close(channelNumbers)
}

func main() {
	fmt.Println("Go!")

	numbersCount := 10
	maxValue := 100

	channelNumbers := make(chan int, numbersCount)
	channelAvg := make(chan float32)
	channelDone := make(chan bool)
	go generateNumbers(numbersCount, maxValue, channelNumbers)
	go calcAverage(channelNumbers, channelAvg)
	go printAverage(channelAvg, channelDone)

	<-channelDone
	fmt.Println("Exit")
}
