package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	count int
	sync.Mutex
}

func serve(wg *sync.WaitGroup, hits *Hits, iteration int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()

	hits.count += 1
	fmt.Println("served iteration", iteration)
}

func main() {

	var wg sync.WaitGroup
	hitsCounter := Hits{}

	for i := 0; i < 20; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitsCounter, iteration)
	}

	fmt.Println("Waiting for goroutines")
	wg.Wait()

	// locking again just to be safe and future proof
	hitsCounter.Lock()
	totalHits := hitsCounter.count
	hitsCounter.Unlock()

	fmt.Println("Total hits", totalHits)
}
