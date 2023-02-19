package main

import "fmt"

func main() {

	var alasAtas, alasBawah, tinggi float32
	var luas float32
	fmt.Print("Masukkan Nilai Panjang Atas Trapesium : ")
	fmt.Scan(&alasAtas)
	fmt.Print("Masukkan Nilai Panjang Bawah Trapesium : ")
	fmt.Scan(&alasBawah)
	fmt.Print("Masukkan Tinggi Trapesium : ")
	fmt.Scan(&tinggi)

	luas = 0.5 * (alasBawah + alasAtas) * tinggi

	fmt.Println("Luas Trapesium adalah ", luas)
}
