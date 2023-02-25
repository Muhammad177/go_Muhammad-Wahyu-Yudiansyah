package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var hasil []int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				hasil = append(hasil, i, j)
				fmt.Print(arr, " Tujuan ", target, " Penjumlahan Dari Array Index : ")
			}
		}
	}
	return hasil
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
}
