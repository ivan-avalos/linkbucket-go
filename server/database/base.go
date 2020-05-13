/*
 *  base.go
 *  Copyright (C) 2020  Iván Ávalos <ivan.avalos.diaz@hotmail.com>
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Affero General Public License as
 *  published by the Free Software Foundation, either version 3 of the
 *  License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU Affero General Public License for more details.
 *
 *  You should have received a copy of the GNU Affero General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

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
	db.Debug().AutoMigrate(&Job{})
}

// DB returns the DB object
func DB() *gorm.DB {
	return db
}
