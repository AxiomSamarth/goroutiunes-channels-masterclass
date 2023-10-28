package main

import (
	"fmt"
	"sync"
	"time"
)

func processOne(wg *sync.WaitGroup, stopCh chan bool) {
	defer wg.Done()

	for {
		select {
		case <-stopCh:
			fmt.Println("Stopping processOne")
			return
		default:
			for i := 0; i < 5; i++ {
				fmt.Println("Running ProcessOne")
				time.Sleep(1 * time.Second)
			}
			time.Sleep(5 * time.Second)
			fmt.Println("Simulating an error. Triggering close of stopChannel")
			close(stopCh)
		}
	}
}

func processTwo(wg *sync.WaitGroup, stopCh chan bool) {
	defer wg.Done()

	for {
		select {
		case <-stopCh:
			fmt.Println("Stopping processTwo")
			return
		default:
			fmt.Println("Running ProcessTwo")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	stopCh := make(chan bool)
	wg.Add(2)

	go processOne(&wg, stopCh)
	go processTwo(&wg, stopCh)

	wg.Wait()
}
