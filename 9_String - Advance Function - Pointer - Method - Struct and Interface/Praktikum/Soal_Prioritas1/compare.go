package main

import (
	"fmt"
	"strings"
)

func compare(a, b string) string {
	if strings.Contains(a, b) {
		return b
	} else if strings.Contains(b, a) {
		return a
	} else {
		return "null"
	}

}

func main() {

	fmt.Println(compare("AKA", "AKASHI"))
	fmt.Println(compare("KANGGONG", "KANG"))
	fmt.Println(compare("KI", "KIJANG"))
	fmt.Println(compare("ILALANG", "ILA"))

}
