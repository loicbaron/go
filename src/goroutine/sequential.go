package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"aaa",
		"https://usertesting.com",
		"https://app.usertesting.com",
		"https://recorder.usertesting.com",
	}
	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Link %s might be down!\n", link)
		return
	}
	fmt.Printf("Link %s is OK\n", link)
	return
}
