package main

import "fmt"

func main() {
	var nilai int32
	var nilai_huruf string
	fmt.Println("**********************************")
	fmt.Println("Menentukan nilai menggunakan Grade")
	fmt.Println("**********************************")
	fmt.Print("Masukkan Nilai Angka : ")
	fmt.Scan(&nilai)

	if nilai >= 80 && nilai <= 100 {
		nilai_huruf = "A"
	} else if nilai >= 80 && nilai <= 100 {
		nilai_huruf = "B"
	} else if nilai >= 50 && nilai <= 64 {
		nilai_huruf = "C"
	} else if nilai >= 35 && nilai <= 49 {
		nilai_huruf = "D"
	} else if nilai >= 0 && nilai <= 34 {
		nilai_huruf = "E"

	}
	fmt.Print("Selamat Nilai Anda termasuk dalam : ", nilai_huruf)
}
