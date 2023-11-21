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

func (bar *Barrier) TryWait(timeout time.Duration) (success bool) {
	blocker := <-bar.blockers
	select {
	case <-blocker:
		return true
	case <-time.After(timeout):
		return false
	}
}

func GetNewBarrier(workersCount int) Barrier {
	barier := Barrier{}
	barier.blockers = blockersGenerator(workersCount)
	return barier
}

func blockersGenerator(capacity int) <-chan chan int {
	out := make(chan chan int)
	go func() {
		for {
			blocker := make(chan int)
			for i := 0; i < capacity; i++ {
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
	sec := getRandomSeconds()
	time.Sleep(time.Duration(sec) * time.Second) // do some work
	fmt.Printf("Work %d in progress, duration: %d sec. Waiting...\n", id, sec)

	if barrier.TryWait(timeout) {
		fmt.Printf("End work %d: success\n", id)
	} else {
		fmt.Printf("End work %d: timeout\n", id)
	}
}

func main() {
	fmt.Println("GO!")

	barierCpacity := 3
	workersCount := 14

	bar := GetNewBarrier(barierCpacity)

	timeout := time.Second * 3

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
