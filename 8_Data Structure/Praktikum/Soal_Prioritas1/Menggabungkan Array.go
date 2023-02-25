package main

import "fmt"

func ArrayMarge(ArrayA, ArrayB []string) []string {
	mergeArray := make(map[string]bool)

	for _, kata := range ArrayA {
		mergeArray[kata] = true
	}
	for _, kata := range ArrayB {
		mergeArray[kata] = true
	}
	ArrayC := []string{}
	for kata := range mergeArray {
		ArrayC = append(ArrayC, kata)
	}
	return ArrayC
}

func main() {
	fmt.Println(ArrayMarge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "devil jin"}))
	fmt.Println(ArrayMarge([]string{}, []string{"eddie", "steve", "devil jin"}))
	fmt.Println(ArrayMarge([]string{}, []string{}))
}
