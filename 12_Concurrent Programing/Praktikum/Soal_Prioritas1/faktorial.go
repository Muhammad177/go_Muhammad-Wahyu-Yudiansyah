package main

import (
	"fmt"
	"sync"
)

func factorial(n int, wg *sync.WaitGroup, mu *sync.Mutex, result *int) {
	defer wg.Done()

	fact := 1
	for i := 1; i <= n; i++ {
		mu.Lock()
		fact *= i
		*result = fact
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var result int

	n := 5 // menghitung faktorial dari 5

	wg.Add(1)
	go factorial(n, &wg, &mu, &result)

	wg.Wait()
	fmt.Printf("Faktorial dari %d adalah %d\n", n, result)
}
