package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	
	// Make a slice of length dy
	p := make([][]uint8, dy)
	
	// iterate over all p indices
	for y := range p {
	
		// place slice in slice of length dx
		p[y] = make([]uint8, dx)
		
		// iterate over slice of length dx
		for x := 0; x < dx; x++ {
			p[y][x] = uint8(x*y)
		}
	}
	// return slice of length dy
	return p
}

func main() {
	pic.Show(Pic)
}
