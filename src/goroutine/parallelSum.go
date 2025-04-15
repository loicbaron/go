package main

import (
	"fmt"
	"sync"
)

func sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func sumChunk(chunk []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	// Calculate the sum of the chunk
	sum := sum(chunk)
	// Send the sum to the channel
	ch <- sum
}

func main() {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	// Step 1: Calculate the sum of the array sequentially
	fmt.Println("sequential sum", sum(arr))

	// Step 2: Calculate the sum of the array in parallel using goroutines
	numGoroutines := 4
	// Create a channel to collect results
	ch := make(chan int, numGoroutines)
	// Use WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	chunkSize := len(arr) / numGoroutines
	fmt.Println("Chunk size:", chunkSize)

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(arr)
		}
		chunk := arr[start:end]
		fmt.Println("Chunk:", chunk)
		// Increment the WaitGroup counter for each goroutine
		wg.Add(1)
		// Create a goroutine for each chunk
		go sumChunk(chunk, ch, &wg)
	}
	// Wait for all goroutines to finish
	wg.Wait()

	// Close the channel
	close(ch)

	// Collect the results from the channel
	totalSum := 0
	for sum := range ch {
		totalSum += sum
	}

	// Print the final sum
	fmt.Println("Parallel sum:", totalSum)
}
