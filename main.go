package main

import (
	"log"

	"github.com/Kunniii/go_gin_gorm_test/initializers"
	"github.com/gin-gonic/gin"
)

// this function is run before main()
func init() {
	gin.SetMode(gin.ReleaseMode)
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "OK",
		})
	})

	log.Fatal(router.Run())
}
