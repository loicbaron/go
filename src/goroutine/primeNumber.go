package main

import (
	"fmt"
	"sync"
)

/*
Challenge: Concurrent Prime Number Calculation

Problem Description:
Write a Go program that finds all prime numbers up to a given number n concurrently.
Use channels and goroutines to calculate the primes in parallel.

Requirements:
The program should calculate all prime numbers up to n.

Split the work across multiple goroutines. Each goroutine should handle a different range of numbers.

Use a channel to send the prime numbers found by each goroutine back to the main goroutine.

Use sync.WaitGroup to synchronize the completion of all goroutines.

The main goroutine should collect the results from the channels and print all prime numbers.

Example:
Input: 30
Output: 2, 3, 5, 7, 11, 13, 17, 19, 23, 29

Constraints:
n can be large (e.g., 10,000 or more).

Efficient prime number checking and good use of goroutines and channels are important.

Hint:
A prime number is only divisible by 1 and itself.
You can divide the task by splitting the range from 2 to n into several smaller chunks
and assign each chunk to a goroutine.
*/

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeChunk(chunk []int, ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	// Calculate prime numbers of the chunk
	fmt.Println("Chunk:", chunk)
	output := []int{}
	for i := range chunk {
			if(isPrime(chunk[i])) {
				output = append(output, chunk[i])
			}
	}
	// Send the output to the channel
	ch <- output
}

func main() {
	input := 30 // range
	output := []int{}
	// for i := range input {
	// 		if(isPrime(i)) {
	// 			output = append(output, i)
	// 		}
	// }
	// fmt.Println("Sequential: prime numbers up to", input, "are:", output)

	// Parallel calculation using goroutines
	numGoroutines := 4
	// Create a channel to collect results
	ch := make(chan []int, numGoroutines)
	// Use WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	chunkSize := input / numGoroutines
	fmt.Println("Chunk size:", chunkSize)

	arr := make([]int, input)
	for i := range arr { arr[i] = i }

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == numGoroutines-1 {
			end = len(arr)
		}
		chunk := arr[start:end]
		// Increment the WaitGroup counter for each goroutine
		wg.Add(1)
		// Create a goroutine for each chunk
		go isPrimeChunk(chunk, ch, &wg)
	}
	// Wait for all goroutines to finish
	wg.Wait()

	// Close the channel
	close(ch)

	for numbers := range ch {
		output = append(output, numbers...)
	}
	fmt.Println("Parallel: prime numbers up to", input, "are:", output)
}
