package main

import (
<<<<<<< HEAD
	"backend/config"
	"backend/routes"
=======
	"backend/routers"
>>>>>>> 54962e854190a18c51ae3ee07cac22c5c8e940dc
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8080")

<<<<<<< HEAD
}
=======
// func init() {
// 	config.IntialDatabase()
// }
>>>>>>> 54962e854190a18c51ae3ee07cac22c5c8e940dc
