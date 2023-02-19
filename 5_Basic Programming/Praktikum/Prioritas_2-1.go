package main

import "fmt"

func main() {
	var jumlahbintang int32
	fmt.Println("Tutorial Membuat Tabel bintang ")
	fmt.Println("================================")
	fmt.Print("Silahkan masukkan jumlah baris : ")
	fmt.Scan(&jumlahbintang)

	for baris := 1; baris <= int(jumlahbintang); baris++ {

		for spasi := 20; spasi >= baris; spasi-- {
			fmt.Printf(" ")
		}

		for kolom := 1; kolom <= baris; kolom++ {
			fmt.Printf("*")

		}
		for kolom := (baris - 1); kolom >= 1; kolom-- {
			fmt.Printf("*")

		}
		fmt.Printf("\n")
	}
}
