package main

import "basic_trade/database"

func main() {
	database.InitDB()
	database.GetDB()
}
