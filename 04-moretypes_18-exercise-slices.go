package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	vals := [][]uint8{}
	for y := 0; y < dy; y++ {
		vals = append(vals, make([]uint8, dx))
		for x := 0; x < dy; x++ {
			// vals[y][x] = uint8((x + y) / 2)
			// vals[y][x] = uint8(x * y)
			vals[y][x] = uint8(x ^ y)
		}
	}
	return vals
}

func main() {
	pic.Show(Pic)
}
