package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	amount = 100
	mu   = sync.Mutex{}
	wg     = sync.WaitGroup{}
)

func stingy() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		amount += 10
		mu.Unlock()

		time.Sleep(time.Millisecond)
	}

	fmt.Println("Stingy Done")
}

func spendy() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		amount -= 10
		mu.Unlock()

		time.Sleep(time.Millisecond)
	}

	fmt.Println("Spendy Done")
}

func main() {
	wg.Add(2)

	go stingy()
	go spendy()

	wg.Wait()

	fmt.Println(amount)
}