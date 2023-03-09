package main

import "fmt"

func toRomanNumeral(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""

	for i := 0; i < len(vals); i++ {

		for num >= vals[i] {
			roman += syms[i]
			num -= vals[i]
		}
	}

	return roman
}

func main() {
	// contoh penggunaan
	fmt.Println(toRomanNumeral(4))    // output: IV
	fmt.Println(toRomanNumeral(9))    // output: IX
	fmt.Println(toRomanNumeral(23))   // output: XXIII
	fmt.Println(toRomanNumeral(2021)) // output: MMXXI
	fmt.Println(toRomanNumeral(1646)) // output: MDCXLVI
}
