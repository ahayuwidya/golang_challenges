package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var gelato_db *sql.DB

type Products struct {
	Id          int
	ProductName string
	CreatedAt   string
	UpdatedAt   string
}

type Variants struct {
	Id          int
	VariantName string
	Quantity    int
	ProductId   int
	CreatedAt   string
	UpdatedAt   string
}

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "gelato_store",
	}

	// Get a database handle.
	gelato_db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := gelato_db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Database connected! \n")

	// createProduct("Cocomelon Test 2")
	// updateProduct("Cocomelon LATEST UPDATE", "Cocomelon")
	// getProductById(73)
	createVariant("Choco Cheese", 50, 2)
	// updateVariantById(1, 96)
	// deleteVariantById(8)
	// getProductWithVariant("Matcha Almond")
}

func createProduct(productname string) {
	var product = Products{}

	sqlStatement := `INSERT INTO products (product_name) VALUES (?)`
	result, err := gelato_db.Exec(sqlStatement, productname)
	if err != nil {
		fmt.Println("Duplicate product name!")
		// panic(err)
	} else {
		lastInsertID, err := result.LastInsertId()
		if err != nil {
			fmt.Println("err2")
			panic(err)
		}
		sqlRetrieve := `SELECT * FROM products WHERE id = ?`
		err = gelato_db.QueryRow(sqlRetrieve, lastInsertID).Scan(&product.Id, &product.ProductName, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			fmt.Println("err3")
			panic(err)
		}
		fmt.Printf("Successfully added ID: %d, Product Name: %s, Created At: %s, Updated At: %s\n", product.Id, product.ProductName, product.CreatedAt, product.UpdatedAt)
	}
}

func updateProduct(oldproductname, newproductname string) {
	var product = Products{}

	sqlStatement := `UPDATE products SET product_name = ?, updated_at = CURRENT_TIMESTAMP WHERE product_name = ?`
	_, err := gelato_db.Exec(sqlStatement, newproductname, oldproductname)
	if err != nil {
		fmt.Println("err1")
		panic(err)
	}

	sqlRetrieve := `SELECT * FROM products ORDER BY updated_at DESC LIMIT 1`
	err = gelato_db.QueryRow(sqlRetrieve).Scan(&product.Id, &product.ProductName, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		fmt.Println("err2")
		panic(err)
	}
	fmt.Printf("Successfully updated: %s -- ID: %d, New Product Name: %s, Updated At: %s\n", oldproductname, product.Id, product.ProductName, product.UpdatedAt)
}

func getProductById(id int) {
	var product = Products{}

	sqlRetrieve := `SELECT * FROM products WHERE id = ?`
	err := gelato_db.QueryRow(sqlRetrieve, id).Scan(&product.Id, &product.ProductName, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		fmt.Println("Product ID not found")
		// panic(err)
	} else {
		fmt.Printf("ID: %d, Product Name: %s\n", product.Id, product.ProductName)
	}
}

func createVariant(varname string, varquantity, productid int) {
	var variant = Variants{}

	sqlStatement := `INSERT INTO variants (variant_name, quantity, product_id) VALUES (?, ?, ?)`
	result, err := gelato_db.Exec(sqlStatement, varname, varquantity, productid)
	if err != nil {
		fmt.Println("Duplicate variant name!")
		// panic(err)
	} else {
		lastInsertID, err := result.LastInsertId()
		if err != nil {
			fmt.Println("err2")
			panic(err)
		}
		sqlRetrieve := `SELECT * FROM variants WHERE id = ?`
		err = gelato_db.QueryRow(sqlRetrieve, lastInsertID).Scan(&variant.Id, &variant.VariantName, &variant.Quantity, &variant.ProductId, &variant.CreatedAt, &variant.UpdatedAt)
		if err != nil {
			fmt.Println("err3")
			panic(err)
		}
		fmt.Printf("Successfully created new variant! \nID: %d, Variant Name: %s, Quantity: %d, Product ID: %d, Created At: %s, Updated At: %s \n", variant.Id, variant.VariantName, variant.Quantity, variant.ProductId, variant.CreatedAt, variant.UpdatedAt)
	}
}

func updateVariantById(variantid, quantity int) {
	var variant = Variants{}

	sqlStatement := `UPDATE variants SET quantity = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?` // quantity, product_name
	_, err := gelato_db.Exec(sqlStatement, quantity, variantid)
	if err != nil {
		fmt.Println("err1")
		panic(err)
	}

	sqlRetrieve := `SELECT * FROM variants ORDER BY updated_at DESC LIMIT 1`
	err = gelato_db.QueryRow(sqlRetrieve).Scan(&variant.Id, &variant.VariantName, &variant.Quantity, &variant.ProductId, &variant.CreatedAt, &variant.UpdatedAt)
	if err != nil {
		fmt.Println("err2")
		panic(err)
	}
	fmt.Printf("Successfully updated variant quantity! \nID: %d, Variant Name: %s, Quantity: %d, Product ID: %d, Created At: %s, Updated At: %s \n", variant.Id, variant.VariantName, variant.Quantity, variant.ProductId, variant.CreatedAt, variant.UpdatedAt)

}

func deleteVariantById(variantid int) {
	var variant = Variants{}

	sqlRetrieve := `SELECT * FROM variants WHERE id = ?`
	err := gelato_db.QueryRow(sqlRetrieve, variantid).Scan(&variant.Id, &variant.VariantName, &variant.Quantity, &variant.ProductId, &variant.CreatedAt, &variant.UpdatedAt)
	if err != nil {
		fmt.Println("Variant ID not found")
		// panic(err)
	} else {
		sqlStatement := `DELETE FROM variants WHERE id = ?`
		_, err = gelato_db.Exec(sqlStatement, variantid)
		if err != nil {
			fmt.Println("err1")
			// panic(err)
		}
		fmt.Printf("Successfully deleted variant! \nID: %d, Variant Name: %s, Product ID: %d \n", variant.Id, variant.VariantName, variant.ProductId)
	}
}

func getProductWithVariant(variantname string) {
	var variant = Variants{}
	var product = Products{}

	sqlRetrieve := `SELECT variants.variant_name, products.product_name FROM products JOIN variants ON variants.product_id = products.id WHERE variants.variant_name = ?`
	err := gelato_db.QueryRow(sqlRetrieve, variantname).Scan(&variant.VariantName, &product.ProductName)
	if err != nil {
		fmt.Println("err1")
		panic(err)
	}
	fmt.Printf("Variant Name: %s, Product Name: %s \n", variant.VariantName, product.ProductName)
}
