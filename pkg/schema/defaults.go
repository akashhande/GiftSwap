package schema

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Init Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../gorm.db")
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	db.DB().SetMaxIdleConns(10)
	db.AutoMigrate(&FamilyMember{})
	db.AutoMigrate(&GiftExchange{})
	DB = db
	return DB
}

// TestDBInit This function will create a temporarily database for running testing cases
func TestDBInit() *gorm.DB {
	test_db, err := gorm.Open("sqlite3", "./../gorm_test.db")
	if err != nil {
		fmt.Println("db err: (TestDBInit) ", err)
	}
	// Migrate schema for testing database
	if err := test_db.AutoMigrate(FamilyMember{}); err != nil {
		fmt.Println("db migration err: (TestDBInit) ", err)
	}
	test_db.DB().SetMaxIdleConns(3)
	test_db.LogMode(true)
	DB = test_db
	return DB
}

// GetDB Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
