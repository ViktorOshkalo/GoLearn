package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numToGuess := rand.Intn(10)
	fmt.Printf("You shoudn't see it. Number: %d\n", numToGuess)

	tries := 3
	success := false
	for i := 0; i < tries; i++ {
		var input int
		fmt.Scan(&input)

		if input == numToGuess {
			fmt.Println("You winner!")
			success = true
			break
		} else if input < numToGuess {
			fmt.Println("Bigger!")
		} else {
			fmt.Println("Lesser!")
		}
	}

	if !success {
		fmt.Println("You looser!")
	}
}
