package main

import "golang.org/x/tour/pic"

// Pic creates a Slice of uint8
func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	line := make([]uint8, dx)
	for y := range img {
		for x := range line {
			line[x] = uint8(x ^ y)
		}
		img[y] = line
	}
	return img
}

func main() {
	pic.Show(Pic)
}
