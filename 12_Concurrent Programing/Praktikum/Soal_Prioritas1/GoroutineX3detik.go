package main

import (
	"fmt"
	"time"
)

func KelipatanBilangan(x int) {
	for i := 1; ; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", i)
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go KelipatanBilangan(4)
	var input string
	fmt.Scanln(&input)
}
