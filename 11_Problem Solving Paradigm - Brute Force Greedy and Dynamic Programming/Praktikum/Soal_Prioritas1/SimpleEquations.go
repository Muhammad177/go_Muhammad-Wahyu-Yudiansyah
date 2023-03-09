package main

import "fmt"

func SimpleEquantions(a, b, c int) {
	for x := 1; x <= a-2; x++ {
		for y := x + 1; y <= a-x-1; y++ {
			z := a - x - y
			if x*y*z == b && x*x+y*y+z*z == c {
				fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
				return
			}
		}
	}
	fmt.Println("Tidak ada solusi yang ditemukan")
}

func main() {
	SimpleEquantions(1, 2, 3)  // no soluation
	SimpleEquantions(6, 6, 14) // 1 2 3
}
