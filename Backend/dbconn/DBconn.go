package dbconn

import (
	"DynamicQRBackend/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Zeeshan@123"
	dbname   = "dynamic-qr-code"
)

var dsn string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Karachi", host, port, user, password, dbname)
var DB *gorm.DB

func DBconn() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	db.AutoMigrate(&models.QRCode{}, &models.QRCodeAnalytics{}, &models.User{})
}
