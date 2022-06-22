package main

import (
	"backend/config"
	"backend/routers"
)

func main() {
	e := routers.Router()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

func init() {
	config.IntialDatabase()
}
