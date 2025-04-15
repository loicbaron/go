package main

import (
	"fmt"
	"math"
	"sort"
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
	// We only need to check up to the square root of the number
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeChunk(start, end int, primes *[]int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Local slice to collect primes for this chunk
	localPrimes := make([]int, 0, (end-start)/2)

	for num := start; num < end; num++ {
		if isPrime(num) {
			localPrimes = append(localPrimes, num)
		}
	}

	// Use a mutex to safely append to the shared primes slice
	mu.Lock()
	*primes = append(*primes, localPrimes...)
	mu.Unlock()
}

var mu sync.Mutex // Mutex for safe concurrent access to primes slice

func main() {
	input := 30 // range
	primes := make([]int, 0, input/2) // Pre-allocate space for prime numbers
	numGoroutines := 4
	var wg sync.WaitGroup

	chunkSize := input / numGoroutines

	// Parallel calculation using goroutines
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = input
		}

		wg.Add(1)
		go isPrimeChunk(start, end, &primes, &wg)
	}

	wg.Wait()

	// Sort the primes slice (optional)
	sort.Ints(primes)

	fmt.Println("Parallel Optimized: prime numbers up to", input, "are:", primes)
}
