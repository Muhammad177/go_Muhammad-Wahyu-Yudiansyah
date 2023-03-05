package main

import (
	"fmt"
	"strconv"
)

type car struct {
	tipe   string
	fuelln float64
}

func (e car) hitungjarak() string {
	var i float64 = e.fuelln / 1.5
	fmt.Print("tipe Mobil ", e.tipe)
	fmt.Printf(" Jarak tempuh dalam Mill dengan bensin %.f Adalah ", e.fuelln)
	var str = strconv.FormatFloat(i, 'f', 2, 64)
	return str + " Mill"
}

func main() {
	var e = car{
		tipe:   "ayla",
		fuelln: 60,
	}
	var c = car{
		tipe:   "datsun",
		fuelln: 50,
	}
	fmt.Println(c.hitungjarak())
	fmt.Println(e.hitungjarak())
}
