package main

import (
	"fmt"
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

func main() {
	fmt.Println("hello")
	input := 30 // range
	output := []int{}
	for i := range input {
			if(isPrime(i)) {
				output = append(output, i)
			}
	}
	fmt.Println("Prime numbers up to", input, "are:", output)
}
