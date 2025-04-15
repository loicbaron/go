/*
Challenge: Custom Sort by Frequency
Problem Description:
Write a Go program that takes an array of integers and sorts the elements by their frequency of occurrence.
If two elements have the same frequency, sort them by their value in ascending order.

Requirements:
Input: A slice of integers.
Output: A new slice of integers sorted by the frequency of their occurrences,
with ties broken by the integer value in ascending order.

Constraints:
You cannot use Go’s built-in sort function directly to solve the problem.
You should implement your own sorting mechanism, utilizing frequency counting and sorting by custom rules.

Example:
Input: arr := []int{4, 3, 1, 6, 4, 1, 3, 4}
Output: [4 4 4 1 1 3 3 6]

Explanation:
The number 4 appears 3 times, so it comes first.
The number 1 appears 2 times, so it comes after 4.
The number 3 also appears 2 times, but it comes after 1 because 1 is smaller than 3.
The number 6 appears 1 time, so it comes last.

Bonus:
Implement this in an efficient way with time complexity better than O(n²).
Consider using a map to count frequencies and a custom sorting mechanism.
*/

package main
import (
	"fmt"
	"sort"
)

type Frequency struct {
	value    int
	frequency int
}

func sortByFrequency(input []int) []int {
	output := []int{}
	// Step 1: Count the frequency of each number
	frequencyMap := make(map[int]int)
	for _, num := range input {
		frequencyMap[num]++
	}
	fmt.Println("Frequency Map:", frequencyMap)

	// Step 2: Convert the frequency map to a slice of Frequency structs
	var freqSlice []Frequency
	for value, freq := range frequencyMap {
		freqSlice = append(freqSlice, Frequency{value: value, frequency: freq})
	}
	fmt.Println("freqSlice", freqSlice)

	// Step 3: Sort the slice
	sort.Slice(freqSlice, func(i, j int) bool {
		// Sort by frequency (descending) and value (ascending)
		if freqSlice[i].frequency == freqSlice[j].frequency {
			return freqSlice[i].value < freqSlice[j].value
		}
		return freqSlice[i].frequency > freqSlice[j].frequency
	})

	// Step 4: Rebuild the sorted output slice
	for _, freq := range freqSlice {
		for i := 0; i < freq.frequency; i++ {
			output = append(output, freq.value)
		}
	}
	return output
}

func main() {
	input := []int{4, 3, 1, 6, 4, 1, 3, 4}
	sortedArr := sortByFrequency(input)
	fmt.Println(sortedArr) // Expected output: [4 4 4 1 1 3 3 6]
}
