package main

import "fmt"

func main() {
	var bilangan int32

	fmt.Println("Penghasil Faktor Suatu Bilangan")
	fmt.Println("================================")
	fmt.Print("Masukkan Bilangan Yang Akan Difaktorkan : ")
	fmt.Scan(&bilangan)

	fmt.Println("Faktor dari bilangan ", bilangan)
	for i := 1; i <= int(bilangan); i++ {
		if int(bilangan)%i == 0 {
			fmt.Println(i)
		}

	}

}
