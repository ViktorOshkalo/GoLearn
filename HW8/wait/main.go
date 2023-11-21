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

func GetNewBarrier(workersCount int) Barrier {
	barier := Barrier{}
	barier.blockers = blockersGenerator(workersCount)
	return barier
}

func blockersGenerator(capacity int) <-chan chan bool {
	out := make(chan chan bool)
	go func() {
		for {
			blocker := make(chan bool)
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

func workerWait(barrier Barrier, id int) {
	fmt.Printf("Start work %d\n", id)

	sec := getRandomSeconds()
	time.Sleep(time.Duration(sec) * time.Second) // do some work
	fmt.Printf("Work %d in progress, duration: %d sec. Waiting...\n", id, sec)

	barrier.Wait()
	fmt.Printf("End work %d\n", id)
}

func main() {
	fmt.Println("GO!")

	barierCpacity := 3
	workersCount := 12
	if (workersCount % barierCpacity) != 0 {
		panic("Workers count must be multiple to barier capacity. Otherwise deadlock")
	}

	bar := GetNewBarrier(barierCpacity)

	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		id := i
		go func() {
			defer wg.Done()
			workerWait(bar, id)
		}()
	}

	wg.Wait()
	fmt.Println("Exit")
}
