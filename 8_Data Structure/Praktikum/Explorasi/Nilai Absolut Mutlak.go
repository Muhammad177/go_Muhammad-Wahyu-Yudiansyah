package main

import (
	"fmt"
	"math"
)

func main() {
	hasildiagonal1 := 0
	hasildiagonal2 := 0
	matriks := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				hasildiagonal1 += matriks[i][j]
			}
			if i+j == 2 {
				hasildiagonal2 += matriks[i][j]
			}
		}
	}
	hasil := hasildiagonal1 - hasildiagonal2
	x := math.Abs(float64(hasil))

	fmt.Print(x)
}
