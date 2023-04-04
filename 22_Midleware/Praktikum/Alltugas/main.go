package main

import (
	"go_Muhammad-Wahyu-Yudiansyah/22_Midleware/Praktikum/Alltugas/database"
	"go_Muhammad-Wahyu-Yudiansyah/22_Midleware/Praktikum/Alltugas/routes"
)

func main() {
	database.InitDB()
	e := routes.New()
	e.Start("8000")

}
