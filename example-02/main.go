package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	queue []int
	lock  chan bool
	wg    sync.WaitGroup
}

func (q *Queue) Enqueue(numbers []int) {
	q.wg.Add(1)
	q.lock <- true
	defer func() {
		q.wg.Done()
		<-q.lock
	}()

	q.queue = append(q.queue, numbers...)
}

func main() {
	queue := &Queue{
		wg:   sync.WaitGroup{},
		lock: make(chan bool, 1),
	}

	for i := 0; i < 100000; i++ {
		go queue.Enqueue([]int{i, i * 2, i * 3})
	}

	queue.wg.Wait()

	// the length of the queue should be 3 * 1,00,000 = 3,00,000
	fmt.Printf("The length of queue is %d\n", len(queue.queue))
}
