package main

import (
	"fmt"
	"regexp"
	"strings"
)

var nonAlphaRegex = regexp.MustCompile(`[^\p{L}]+`)

func clearString(str string) string {
    return nonAlphaRegex.ReplaceAllString(str, "")
}

func reverseString(s string) string {
		result := ""
		for i, _ := range s {
			result = result + string(s[len(s) - i - 1])
		}
    return result
}

func isPalindrome(s string) bool {
		x := clearString(strings.ToLower(s))
		return x == reverseString(x)
}

func main() {
    testStr := "A man, a plan, a canal: Panama"
    result := isPalindrome(testStr)
    fmt.Println(result)  // Expected output: true
}
