package main

import "fmt"

func withoutValue() int {
	var age int // variable declaration
	return age
}

func withType() int {
	var age int = 10 // variable declaration with initial value
	return age
}

func withTypeInferred() int {
	var age = 20 // type will be inferred
	return age
}

func withShortAssign() int {
	age := 30
	return age
}

func main() {
	fmt.Println("My age is", withoutValue())
	fmt.Println("My age is", withType())
	fmt.Println("My age is", withTypeInferred())
	fmt.Println("My age is", withShortAssign())
}
