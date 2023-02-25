package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculsekali(angka string) []int {

	cek := make(map[string]int)
	pisah := strings.Split(angka, "")

	for _, kata := range pisah {
		cek[kata]++
	}
	var hitung []int
	for _, kata := range pisah {
		if cek[kata] == 1 {
			konversi, _ := strconv.Atoi(kata)
			hitung = append(hitung, konversi)
		}
	}
	return hitung
}

func main() {
	fmt.Println(munculsekali("1234123"))
	fmt.Println(munculsekali("76523752"))
	fmt.Println(munculsekali("12345"))
	fmt.Println(munculsekali("1122334455"))
}
