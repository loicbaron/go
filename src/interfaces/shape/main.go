package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	tr := triangle{10, 10}
	sq := square{10}

	printArea(tr)
	printArea(sq)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
