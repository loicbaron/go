package main

import "fmt"

func areAnagrams(str1, str2 string) bool {
    if len(str1) != len(str2) {
			return false
		}
		charCount := make(map[rune]int)
		for _, char := range str1 {
			charCount[char]++
		}
		charCount2 := make(map[rune]int)
		for _, char := range str2 {
			charCount2[char]++
		}
		if len(charCount) != len(charCount2) {
			return false
		}
		for char, count := range charCount {
			if charCount2[char] != count {
				return false
			}
		}
    return true
}

func areAnagramsOptimized(str1, str2 string) bool {
    if len(str1) != len(str2) {
        return false
    }

    charCount := make(map[rune]int)

    // Increment character counts for str1
    for _, char := range str1 {
        charCount[char]++
    }

    // Decrement character counts for str2
    for _, char := range str2 {
        charCount[char]--
    }

    // If all values are zero, then it's an anagram
    for _, count := range charCount {
        if count != 0 {
            return false
        }
    }

    return true
}

func main() {
    str1 := "listen"
    str2 := "silent"
    result := areAnagrams(str1, str2)
    fmt.Println(result)  // Expected output: true

		str1 = "triangle"
    str2 = "integral"
    result = areAnagramsOptimized(str1, str2)
    fmt.Println(result)  // Expected output: true

		str1 = "apple"
		str2 = "pale"
		result = areAnagrams(str1, str2)
		fmt.Println(result)  // Expected output: false

		str1 = "apple"
		str2 = "pale"
		result = areAnagramsOptimized(str1, str2)
		fmt.Println(result)  // Expected output: false
}
// Test cases
// 1. "listen" and "silent" should return true
// 2. "triangle" and "integral" should return true
// 3. "apple" and "pale" should return false
