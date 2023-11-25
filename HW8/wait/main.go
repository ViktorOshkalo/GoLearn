package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Barrier struct {
	blockers <-chan chan bool
}

func (bar *Barrier) Wait() {
	blocker := <-bar.blockers
	<-blocker
}

func GetNewBarrier(workersCount int, capacity int) Barrier {
	barier := Barrier{}
	barier.blockers = getNewBlockersGenerator(workersCount, capacity)
	return barier
}

func getNewBlockersGenerator(blockersCount int, capacity int) <-chan chan bool {
	out := make(chan chan bool)
	go func() {
		for blockersSent := 0; blockersSent < blockersCount; {
			blocker := make(chan bool)
			for j := 0; j < capacity; j++ {
				out <- blocker
				blockersSent++
			}
			close(blocker)
		}
		fmt.Println("Blockers generator finished.")
	}()
	return out
}

func getRandomSeconds() int32 {
	return rand.Int31n(15)
}

func runWorker(barrier Barrier, workId int) {
	fmt.Printf("Start work %d\n", workId)

	sec := getRandomSeconds()
	time.Sleep(time.Duration(sec) * time.Second) // do some work
	fmt.Printf("Work %d in progress, duration: %d sec. Waiting...\n", workId, sec)

	barrier.Wait()
	fmt.Printf("End work %d\n", workId)
}

func main() {
	fmt.Println("GO!")

	barierCpacity := 3
	workersCount := 12
	if (workersCount % barierCpacity) != 0 {
		panic("Workers count must be divisible for barier capacity. Otherwise deadlock")
	}

	barier := GetNewBarrier(workersCount, barierCpacity)

	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		workId := i
		go func() {
			defer wg.Done()
			runWorker(barier, workId)
		}()
	}

	wg.Wait()
	fmt.Println("Exit")
}
