package main

import "fmt"

func jadikanbinner(n int) []string {
	biner := make([]string, n+1)
	for i := 0; i <= n; i++ {
		biner[i] = fmt.Sprintf("%b", i)
	}
	return biner
}

func main() {
	fmt.Println(jadikanbinner(2))
	fmt.Println(jadikanbinner(5))
}
