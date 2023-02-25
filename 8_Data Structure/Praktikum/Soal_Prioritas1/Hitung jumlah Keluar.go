package main

import "fmt"

func Mapping(slice []string) map[string]int {
	mergeArray := make(map[string]int)
	for i, kata := range slice {
		if kata == slice[i] {
			mergeArray[kata] += 1
		}
	}
	return mergeArray
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "asd", "qwe"}))
	fmt.Println(Mapping([]string{}))
}
