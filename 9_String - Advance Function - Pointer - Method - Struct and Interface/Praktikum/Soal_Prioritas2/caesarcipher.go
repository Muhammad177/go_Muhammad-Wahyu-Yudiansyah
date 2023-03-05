package main

import "fmt"

func caesar(offset int, input string) string {
	var result string

	for _, abjad := range input {
		if abjad >= 'a' && abjad <= 'z' {
			result += string('a' + (abjad-'a'+rune(offset))%26)
		} else if abjad >= 'A' && abjad <= 'Z' {
			result += string('A' + (abjad-'A'+rune(offset))%26)
		} else {
			result += string(abjad)
		}
	}

	return result
}

func main() {
	var c string
	var d, e int
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
	fmt.Print("Masukkan Perulangan : ")
	fmt.Scan(&e)
	for i := 0; i < e; i++ {
		fmt.Print("Masukkan Loncatan :")
		fmt.Scan(&d)
		fmt.Print("Masukkan Huruf :")
		fmt.Scan(&c)
		fmt.Println(caesar(d, c))
	}

}
