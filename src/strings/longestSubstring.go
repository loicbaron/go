package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[rune]int)
		maxLength := 0
		start := 0
		for i, char := range s {
			fmt.Println(i, char, string(char), charIndex)
			if index, found := charIndex[char]; found && index >= start {
				start = index + 1
			}
			charIndex[char] = i
			if length := i - start + 1; length > maxLength {
				maxLength = length
			}
		}
    return maxLength
}

func main() {
    str := "abcabcbb"
    result := lengthOfLongestSubstring(str)
    fmt.Println(result)  // Expected output: 3
}
