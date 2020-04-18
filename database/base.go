package database

import (
	"fmt"
	"log"
	"os"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Init initialises DB
func Init() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbURI := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName)
	log.Print(dbURI)

	conn, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	db = conn
	if os.Getenv("DEBUG_MODE") == "true" {
		db.LogMode(true)
	}
	db.Debug().AutoMigrate(&User{})
	db.Debug().AutoMigrate(&Link{})
	db.Debug().AutoMigrate(&PasswordReset{})
	db.Debug().AutoMigrate(&Tag{})
}

// DB returns the DB object
func DB() *gorm.DB {
	return db
}
