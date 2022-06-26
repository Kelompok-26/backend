package main

import (
	"backend/config"
	"backend/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8080")

}
