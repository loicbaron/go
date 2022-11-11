package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := os.Args[1]
	path, err := os.Getwd()
	check(err)
	file := fmt.Sprintf("%s/%s", path, filename)
	fmt.Printf("Open filename %s\n", file)
	dat, err := os.ReadFile(file)
	check(err)
	fmt.Print(string(dat))
}
