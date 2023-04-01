package main

import (
	"go_Muhammad-Wahyu-Yudiansyah/21_ORM-MVC/Praktikum/prioritas1/database"
	"go_Muhammad-Wahyu-Yudiansyah/21_ORM-MVC/Praktikum/prioritas1/routes"
)

func main() {
	database.InitDB()
	e := routes.New()
	e.Start("8000")

}
