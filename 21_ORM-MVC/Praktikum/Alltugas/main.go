package main

import (
	"go_Muhammad-Wahyu-Yudiansyah/21_ORM-MVC/Praktikum/Alltugas/database"
	"go_Muhammad-Wahyu-Yudiansyah/21_ORM-MVC/Praktikum/Alltugas/routes"
)

func main() {
	database.InitDB()
	e := routes.New()
	e.Start("8000")

}
