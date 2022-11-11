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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Link %s might be down!\n", link)
		c <- "down"
		return
	}
	fmt.Printf("Link %s is OK\n", link)
	c <- "up"
	return
}
