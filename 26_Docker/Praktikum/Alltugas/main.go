package main

import (
	"go_Muhammad-Wahyu-Yudiansyah/26_Docker/Praktikum/Alltugas/database"
	"go_Muhammad-Wahyu-Yudiansyah/26_Docker/Praktikum/Alltugas/routes"
)

func main() {

	database.InitDB()
	e := routes.New()
	e.Start("8000")

}
