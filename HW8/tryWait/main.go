package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Barrier struct {
	capacity             int
	blocker              chan bool
	awaitingWorkersCount int
	mutex                sync.Mutex
}

func (barier *Barrier) takeBlocker() chan bool {
	barier.mutex.Lock()
	blocker := barier.blocker
	barier.awaitingWorkersCount++
	if barier.awaitingWorkersCount == barier.capacity {
		close(barier.blocker)
		barier.awaitingWorkersCount = 0
		barier.blocker = make(chan bool)
	}
	barier.mutex.Unlock()
	return blocker
}

func (bar *Barrier) tryDiscardBlocker() (success bool) {
	bar.mutex.Lock()
	success = false
	isAllAvailableBlockersTakenBeforeTimeout := bar.awaitingWorkersCount == 0
	if !isAllAvailableBlockersTakenBeforeTimeout {
		bar.awaitingWorkersCount--
		success = true
	}
	bar.mutex.Unlock()
	return
}

func (bar *Barrier) TryWait(timeout time.Duration) (success bool) {
	blocker := bar.takeBlocker()
	select {
	case <-blocker:
		return true
	case <-time.After(timeout):
		discardSuccess := bar.tryDiscardBlocker()
		return !discardSuccess
	}
}

func GetNewBarrier(workersCount int, barierCapacity int) *Barrier {
	barier := Barrier{}
	barier.capacity = barierCapacity
	barier.blocker = make(chan bool)
	return &barier
}

func getRandomSeconds() int32 {
	return rand.Int31n(15)
}

func runWorker(barrier *Barrier, id int, timeout time.Duration) {
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
	timeout := time.Second * 3

	barier := GetNewBarrier(workersCount, barierCpacity)
	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		id := i
		go func() {
			defer wg.Done()
			runWorker(barier, id, timeout)
		}()
	}

	wg.Wait()
	fmt.Println("Exit")
}
