package main

import (
	"log"

	"github.com/Kunniii/go_gin_gorm_test/controllers"
	"github.com/Kunniii/go_gin_gorm_test/internal"
	"github.com/gin-gonic/gin"
)

// this function will run before main()
func init() {
	gin.SetMode(gin.ReleaseMode)
	internal.LoadEnv()
	internal.ConnectDB()
}

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.GetAllPosts)
	router.GET("/posts/:id", controllers.GetPostById)
	router.PUT("/posts/:id", controllers.UpdatePost)

	router.DELETE("/posts/:id", controllers.DeletePostById)

	log.Fatal(router.Run())
}
