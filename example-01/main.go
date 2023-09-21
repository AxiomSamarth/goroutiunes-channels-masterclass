package main

import (
	"fmt"
	"math/rand"
	"time"
)

func populateNumbers(numbers []int) {
	for index := 0; index < len(numbers); index++ {
		numbers[index] = rand.Intn(1000)
	}
}

func findSum(numbers []int, sumChannel chan<- int) {
	sumChannel <- findSumIteratively(numbers)
}

func findSumIteratively(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	// create and populate Slice of integers of length 10 million
	numbers := make([]int, 100000000)
	populateNumbers(numbers)

	// create smaller partitions of the Slice and find the sum of each partition concurrently
	partitions := 10
	partitionLength := len(numbers) / partitions
	sumChannel := make(chan int)
	sum := 0

	// gather the results from each partition and sum them up to find the final result
	startTime := time.Now()
	for i := 0; i < len(numbers); i += partitionLength {
		go findSum(numbers[i:i+partitionLength], sumChannel)
	}

	for i := 0; i < partitions; i++ {
		sum += <-sumChannel
	}

	fmt.Printf("Sum of integers calculated concurrently is %d and took %v\n", sum, time.Since(startTime))

	startTime = time.Now()
	fmt.Printf("Sum of integers calculated iteratively is %d and took %v\n", findSumIteratively(numbers), time.Since(startTime))
}
