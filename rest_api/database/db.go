package database

import (
	"fmt"
	"log"
	"os"
	"rest_api/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// connect and migrate to DB
func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	dsn_config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	// connect to DB
	db, err := gorm.Open(mysql.Open(dsn_config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	fmt.Println("Successfully connected to database.")

	// migrate to DB
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Order{})
	fmt.Println("Database migrated successfully.")
	if db == nil {
		log.Fatal("database connection is nil")
	}

	return db
}

// func MigrateDB() {

// 	// migrate to DB
// 	// db.Debug().AutoMigrate(models.Item{}, models.Order{}) // order_db.Debug().AutoMigrate(&models.Item{}, &models.Order{})
// 	// db.Debug().AutoMigrate(models.Order{})

// 	if db == nil {
// 		// log.Fatal("CHECK!!!! Database connection is nil")
// 		fmt.Println("CHECK!!!! Database connection is nil")
// 	}
// 	if db != nil {
// 		// log.Fatal("CHECK!!!! Database connection is NOTTT nil")
// 		fmt.Println("CHECK!!!! Database connection is NOTTT nil")
// 	}
// }

// // access to DB
// func GetDB() *gorm.DB {
// 	if db == nil {
// 		fmt.Println("Database is nil")
// 	} else {
// 		fmt.Println("Succcessfully get database")
// 	}
// 	return db
// }
