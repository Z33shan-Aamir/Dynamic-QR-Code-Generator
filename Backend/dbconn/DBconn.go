package dbconn

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "host.docker.internal"
	port     = 4050
	user     = "postgres"
	password = "ZeesHANAamir@931360"
	dbname   = "movie-data-provider"
)

var dsn string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Karachi", host, port, user, password, dbname)
var DB *gorm.DB

func DBconn() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	db.AutoMigrate()
}
