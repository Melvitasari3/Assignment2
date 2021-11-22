package database

import (
	"assignment2/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "root"
	password = "244"
	dbPort   = "3306"
	dbName   = "orders_by"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, dbPort, dbName)
	dsn := config
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting", err)
	}
	fmt.Println("sukses koneksi ke db")
	err = db.Debug().AutoMigrate(models.Orders{}, models.Item{})
	if err != nil {
		return
	}
}
func GetDB() *gorm.DB {
	return db
}
