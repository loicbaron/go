package main

import "fmt"

func firstNonRepeating(s string) string {
		charCount := make(map[string]int)
		for _, char := range s {
			fmt.Println(string(char))
			charCount[string(char)]++
		}
		fmt.Println(charCount)
		for _, char := range s {
			if charCount[string(char)] == 1 {
				return string(char)
			}
		}
		return ""
}

func main() {
    testStr := "swiss"
    result := firstNonRepeating(testStr)
    fmt.Println(result)  // Expected output: "w"

		testStr = "aabbcc"
    result = firstNonRepeating(testStr)
    fmt.Println(result)  // Expected output: ""
}
