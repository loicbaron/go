package main

import (
	"fmt"
	"net/http"
	"time"
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

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Link %s might be down!\n", link)
		c <- link
		return
	}
	fmt.Printf("Link %s is OK\n", link)
	c <- link
	return
}
