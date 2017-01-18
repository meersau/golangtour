package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	outer := make( [][]uint8, dy )
	for i := 0; i < dy; i++ {
		outer[i] = make( []uint8, dx)
		for j := 0; j < dx ; j++ {
			//e := (i+255)*(j+255)
			//e := i*j
			//e := i^j
			e := (i+j)/2
			//e := 5*i - 10*j
			outer[i][j] = uint8(e)
		}
	}
	return outer 
}

func main() {
	pic.Show(Pic)
}

