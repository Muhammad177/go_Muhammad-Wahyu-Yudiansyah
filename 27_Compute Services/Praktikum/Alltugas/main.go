package main

import (
	"Alltugas/database"
	"Alltugas/routes"
)

func main() {

	database.InitDB()
	e := routes.New()
	e.Start("8000")

}
