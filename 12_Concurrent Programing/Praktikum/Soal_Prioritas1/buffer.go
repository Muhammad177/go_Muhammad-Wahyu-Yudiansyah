package main

import (
	"fmt"
	"time"
)

func printMultiplesOfThree(n int) {
	// membuat buffer channel dengan kapasitas 10
	c := make(chan int, 10)

	// membuat goroutine yang mengirim bilangan kelipatan 3 ke channel c
	go func() {
		for i := 1; ; i++ {
			if i%3 == 0 {
				c <- i
			}
		}
	}()

	// menunggu selama 1 detik dan mencetak nilai-nilai yang diterima dari channel c
	time.Sleep(1 * time.Second)
	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
}

func main() {
	// mencetak 10 bilangan kelipatan 3
	printMultiplesOfThree(10)
}
