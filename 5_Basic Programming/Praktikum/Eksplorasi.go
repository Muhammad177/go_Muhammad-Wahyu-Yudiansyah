package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	fmt.Println("Menentukan Suatu Huruf atau Angka Palindrome")
	fmt.Println("============================================")
	fmt.Print("Silahkan Masukkan Angka atau Huruf : ", input)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	fmt.Println("Captured :", input)
	for i := 0; i < 1; i++ {
		if input[i] != input[len(input)-1-i] {
			fmt.Println(input, " Bukan Palindrome")
		} else {
			fmt.Println(input, " Adalah Palindrome")
		}

	}
}
