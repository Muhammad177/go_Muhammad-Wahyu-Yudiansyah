package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	for _, kata := range s.name {
		if kata >= 'a' && kata <= 'z' {
			nameEncode += string('a' + 'z' - kata)
		} else if kata >= 'A' && kata <= 'Z' {
			nameEncode += string('A' + 'Z' - kata)
		} else {
			nameEncode += string(kata)
		}
	}

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	for _, kata := range s.nameEncode {
		if kata >= 'a' && kata <= 'z' {
			nameDecode += string('a' + 'z' - kata)
		} else if kata >= 'A' && kata <= 'Z' {
			nameDecode += string('A' + 'Z' - kata)
		} else {
			nameDecode += string(kata)
		}
	}

	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + " is : " + c.Decode())
	}
}
