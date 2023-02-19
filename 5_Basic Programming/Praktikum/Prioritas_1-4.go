package main

import "fmt"

func main() {
	var angka int32
	fmt.Println("============================================")
	fmt.Println("Pencetakan Kelipatan 3 Fizz Kelipatan 5 Buzz")
	fmt.Println("============================================")

	fmt.Println("Masukkan Range angka : ")
	fmt.Scan(&angka)

	for i := 1; i <= int(angka); i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz,")
		} else if i%5 == 0 {
			fmt.Print("Buzz,")

		} else if i%3 == 0 {
			fmt.Print("Fizz,")
		} else {
			fmt.Print(i, ",")
		}

	}
}
