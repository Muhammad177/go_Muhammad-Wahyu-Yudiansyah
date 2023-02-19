package main

import "fmt"

func main() {
	var bilangan int32

	fmt.Println("Menentukan Bilangan Ganjil atau Bilangan Genap Golang")
	fmt.Println("=============================================")
	fmt.Print("Silahkan Masukkan Angka untuk ditentukan : ")
	fmt.Scan(&bilangan)

	if bilangan%2 == 0 {
		fmt.Println("Bilangan Merupakan Bilangan Genap")
	} else {
		fmt.Println("Bilangan Merupakan Bilangan Ganjil")
	}
}
