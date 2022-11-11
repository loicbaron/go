package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var nobody person
	loic := person{"Loïc", "Baron", contactInfo{
		"loic@example.com",
		"012345",
	},
	}
	lowic := person{
		firstName: "Low",
		lastName:  "Hic",
		contact:   contactInfo{},
	}

	fmt.Printf("%+v\n", nobody)
	fmt.Println(loic)
	fmt.Println(lowic)
}
