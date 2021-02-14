package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rapando/gin-api/config"
	"github.com/rapando/gin-api/routes"
)

var err error

func main() {
	_ = godotenv.Load()
	config.DB, err = gorm.Open("mysql", os.Getenv("DB_URI"))
	if err != nil {
		log.Printf("unable to connect to db because %v", err)
		os.Exit(3)
	}
	defer config.DB.Close()

	r := routes.Router()
	r.Run()
}
