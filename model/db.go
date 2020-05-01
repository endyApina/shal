package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var conn *gorm.DB

func init() {
	err := godotenv.Load("conf.env")
	if err != nil {
		log.Println(err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	log.Println(dbURI)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Println(err)
		return
	}

	conn = db
	log.Println("Database connection success")

}

//GetDB retrieves the database connection strings
func GetDB() *gorm.DB {
	return conn
}
