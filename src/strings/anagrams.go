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

func main() {
    str1 := "listen"
    str2 := "silent"
    result := areAnagrams(str1, str2)
    fmt.Println(result)  // Expected output: true

		str1 = "triangle"
    str2 = "integral"
    result = areAnagrams(str1, str2)
    fmt.Println(result)  // Expected output: true

		str1 = "apple"
		str2 = "pale"
		result = areAnagrams(str1, str2)
		fmt.Println(result)  // Expected output: false
}
// Test cases
// 1. "listen" and "silent" should return true
// 2. "triangle" and "integral" should return true
// 3. "apple" and "pale" should return false
