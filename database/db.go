package database

import (
	"assignment2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "vita"
	dbPort   = "5432"
	dbName   = "order_db"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",host, user, password, dbName, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
