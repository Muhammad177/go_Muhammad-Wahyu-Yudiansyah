package main

import (
	"fmt"
)

func primeX(number int) int {
	count := 0
	num := 2
	for {
		if isPrime(num) {
			count++
			if count == number {
				return num
			}
		}
		num++
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeX(1)) // output: 2
	fmt.Println(primeX(5)) // output: 11
	fmt.Println(primeX(8)) // output: 19
}
