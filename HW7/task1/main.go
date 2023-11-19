package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printAverage(channelAvg chan float32) {
	fmt.Println("printAvg started")
	avg := <-channelAvg
	fmt.Println("Average value is: ", avg)
}

func calcAverage(channelIn chan int, channelOut chan float32) {
	fmt.Println("calcAverage started")
	var sum int
	var counter int
	for val := range channelIn {
		fmt.Println("recieving from numbers channel: ", val)
		sum += val
		counter++
	}

	avg := float32(sum) / float32(counter)
	fmt.Println("sending average value into channel: ", avg)
	channelOut <- avg
	close(channelOut)
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

	go generateNumbers(numbersCount, maxValue, channelNumbers)
	go calcAverage(channelNumbers, channelAvg)
	go printAverage(channelAvg)

	time.Sleep(time.Second * 2)
	fmt.Println("Exit")
}
