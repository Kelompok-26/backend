package main

import (
	"backend/routers"
)

func main() {
	e := routers.Router()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":3000"))
}

// func init() {
// 	config.IntialDatabase()
// }
