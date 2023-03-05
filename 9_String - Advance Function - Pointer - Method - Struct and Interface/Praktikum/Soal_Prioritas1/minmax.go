package main

import "fmt"

func findMaxMin(numbers []int) (max int, min int) {
	max = numbers[0]
	min = numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}

func main() {
	var a1, a2, a3, a4, a5, a6 int

	numbers := []int{a1, a2, a3, a4, a5, a6}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Masukkan Angka  %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	max, min := findMaxMin(numbers)

	fmt.Printf("Nilai maksimum adalah %d\n", max)
	fmt.Printf("Nilai minimum adalah %d\n", min)

	maxPointer := &max
	minPointer := &min

	fmt.Printf("Max pointer : %v\n", maxPointer)
	fmt.Printf("Nilai dari pointer max : %d\n", *maxPointer)

	fmt.Printf("Min pointer : %v\n", minPointer)
	fmt.Printf("Nilai dari pointer min : %d\n", *minPointer)
}
