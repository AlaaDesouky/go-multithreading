package main

import (
	"sync"
	"time"
)

type Barrier struct {
	total int
	count int
	mutex *sync.Mutex
	cond  *sync.Cond
}

func NewBarrier(size int) *Barrier {
	muToUse := &sync.Mutex{}
	condToUse := sync.NewCond(muToUse)
	return &Barrier{size, size, muToUse, condToUse}
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count -= 1

	if b.count == 0 {
		b.count = b.total
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}

	b.mutex.Unlock()
}

func waitOnBarrier(name string, timeToSleep int, barrier *Barrier) {
	for {
		println(name, "running")
		time.Sleep(time.Duration(timeToSleep) * time.Second)
		println(name, "is waiting on barrier")
		barrier.Wait()
	}
}

func main() {
	barrier := NewBarrier(2)
	go waitOnBarrier("Red", 4, barrier)
	go waitOnBarrier("Blue", 10, barrier)
	time.Sleep(time.Duration(100) * time.Second)
}
