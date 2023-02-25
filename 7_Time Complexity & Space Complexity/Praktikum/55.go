package main

import "fmt"

func pow(x, n int) int {
	total := 1
	for i := 1; i <= n; i++ {
		total = total * x
	}
	return total
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
