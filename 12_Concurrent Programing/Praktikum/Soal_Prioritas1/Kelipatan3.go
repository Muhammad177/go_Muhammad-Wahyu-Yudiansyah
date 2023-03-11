package main

import (
	"fmt"
)

func printKelipatan3(ch chan int) {
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			ch <- i // kirim bilangan kelipatan 3 ke channel
		}
	}
	close(ch) // tutup channel setelah selesai mengirim
}

func main() {
	ch := make(chan int)

	// jalankan goroutine untuk mencetak bilangan kelipatan 3
	go printKelipatan3(ch)

	// baca nilai dari channel dan cetak ke layar
	for val := range ch {
		fmt.Println(val)
	}
}
