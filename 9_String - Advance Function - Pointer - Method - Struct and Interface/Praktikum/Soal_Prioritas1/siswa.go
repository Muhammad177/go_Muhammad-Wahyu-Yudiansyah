package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	Score float64
}

var allStudents = []Student{
	{name: "Rizky", Score: 80},
	{name: "Kobar", Score: 75},
	{name: "Ismail", Score: 70},
	{name: "Umam", Score: 75},
	{name: "Yopan", Score: 60},
}

func (s Student) rata() float64 {
	x := len(allStudents)
	y := float64(x)
	var p float64 = 0
	var result float64 = 0

	for _, student := range allStudents {
		fmt.Println(student.name, " dengan Nilai : ", student.Score)
		p += student.Score
		result += student.Score / y
	}
	sort.Slice(allStudents, func(i, j int) bool {
		return allStudents[i].Score < allStudents[j].Score
	})
	fmt.Println("Nilai Terendah : ", allStudents[0])
	sort.Slice(allStudents, func(i, j int) bool {
		return allStudents[i].Score > allStudents[j].Score
	})
	fmt.Println("Nilai Tertinggi : ", allStudents[0])
	fmt.Println("Jumlah siswa : ", x)
	fmt.Println("Total Nilai : ", p)
	fmt.Print("Nilai rata-ratanya adalah : ")
	return result
}

func main() {

	s := Student{}

	fmt.Print(s.rata())
}
