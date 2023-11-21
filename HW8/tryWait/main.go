package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Barrier struct {
	blockers <-chan chan int
}

func (bar *Barrier) Wait() {
	fmt.Println("Waiting...")
	blocker := <-bar.blockers
	<-blocker
}

func (bar *Barrier) TryWait(timeout time.Duration) (success bool) {
	fmt.Println("Waiting...")
	blocker := <-bar.blockers
	select {
	case <-blocker:
		return true
	case <-time.After(timeout):
		return false
	}
}

func blockersGenerator(length int) <-chan chan int {
	out := make(chan chan int)
	go func() {
		for {
			blocker := make(chan int)
			for i := 0; i < length; i++ {
				out <- blocker
			}
			close(blocker)
		}
	}()
	return out
}

func getRandomSeconds() int32 {
	return rand.Int31n(15)
}

func workerTryWait(barrier Barrier, id int, timeout time.Duration) {
	fmt.Printf("Start work %d\n", id)

	sec := getRandomSeconds()
	time.Sleep(time.Duration(sec) * time.Second) // do some work
	fmt.Printf("Work %d in progress, duration: %d sec\n", id, sec)

	if barrier.TryWait(timeout) {
		fmt.Printf("End work %d\n", id)
	} else {
		fmt.Printf("Timeout on work %d\n", id)
	}
}

func GetNewBarrier(workersCount int) Barrier {
	barier := Barrier{}
	barier.blockers = blockersGenerator(workersCount)
	return barier
}

func main() {
	fmt.Println("GO!")

	barierCpacity := 3
	workersCount := 14

	bar := GetNewBarrier(barierCpacity)

	timeout := time.Second * 5

	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		id := i
		go func() {
			defer wg.Done()
			workerTryWait(bar, id, timeout)
		}()
	}

	wg.Wait()
	fmt.Println("Exit")
}
