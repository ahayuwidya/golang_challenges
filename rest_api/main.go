package main

import (
	"fmt"
	"rest_api/database"
	"rest_api/models"
)

func main() {
	database.InitDB()
	fmt.Println("main 1\n")
	// database.MigrateDB()
	// fmt.Println("main 2\n")
	// database.GetDB()
	// fmt.Println("main 3")
	CreateOrder("test123")
	// fmt.Println("main 3")
}

func CreateOrder(custname string) {
	fmt.Println("hello CreateOrder")
	orderitems_db := database.InitDB()
	if orderitems_db == nil {
		fmt.Println("!!!! orderitems_db Database connection is nil")
	}

	order := models.Order{
		CustomerName: custname,
	}

	err := orderitems_db.Create(&order).Error
	if err != nil {
		fmt.Println("Error creating order data")
		return
	}

	fmt.Println("New order data: ", order)
}

// Make endpoint to :
// Create Order
// URL: http://localhost:8080/orders
// Method: POST
// Body raw json: {"orderedAt":"2019-11-09T21:21:46+00:00","customerName":"Fitri","items":[{"name":"iPhone","description":"iPhone 11 Pro","quantity":2}]}
