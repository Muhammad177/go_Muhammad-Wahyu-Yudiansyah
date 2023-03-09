package main

import "fmt"

func segitigapascal(numRows int) [][]int {
	segitiga := [][]int{}
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = segitiga[i-1][j-1] + segitiga[i-1][j]
		}
		segitiga = append(segitiga, row)
	}
	return segitiga
}

func main() {
	fmt.Println(segitigapascal(5))
}
