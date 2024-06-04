package main

import (
	"rest_api/database"
	"rest_api/routers"
)

func main() {
	database.InitDB()

	var PORT = ":8080"
	routers.StartApp().Run(PORT)
}
