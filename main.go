package main

import (
	"assignment-2/database"
	"assignment-2/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}