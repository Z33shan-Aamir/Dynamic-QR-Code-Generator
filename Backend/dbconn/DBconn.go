package dbconn

import (
	"DynamicQRBackend/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     string = os.Getenv("DB_HOST")
	user     string = os.Getenv("DB_USER")
	password string = os.Getenv("DB_PASSWORD")
	dbname   string = os.Getenv("DB_NAME")
	port     string = os.Getenv("DB_PORT")
)

var dsn string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Karachi", host, port, user, password, dbname)
var DB *gorm.DB

func DBconn() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	db.AutoMigrate(&models.QRCode{}, &models.QRCodeAnalytics{}, &models.User{})
}
