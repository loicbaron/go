package main

import "fmt"

func reverseString(s string) string {
		result := ""
		for i, letter := range s {
			fmt.Println(i, letter, string(s[len(s) - i - 1]))
			result = result + string(s[len(s) - i - 1])
		}
    return result
}

func main() {
    testStr := "hello"
    reversed := reverseString(testStr)
    fmt.Println(reversed)  // Expected output: "olleh"
}
