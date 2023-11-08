package main

import (
	"log"

	"github.com/Kunniii/go_gin_gorm_test/internal"
	"github.com/Kunniii/go_gin_gorm_test/models"
)

func init() {
	internal.LoadEnv()
	internal.ConnectDB()
}

func main() {
	if err := internal.DB.AutoMigrate(&models.Post{}); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database migration successfully!")
	}
}
