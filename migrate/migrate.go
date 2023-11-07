package main

import (
	"log"

	"github.com/Kunniii/go_gin_gorm_test/initializers"
	"github.com/Kunniii/go_gin_gorm_test/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	if err := initializers.DB.AutoMigrate(&models.Post{}); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database migration successfully!")
	}
}
