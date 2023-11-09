package main

import (
	"log"
	"net/http"

	"github.com/Kunniii/go_gin_gorm_test/controllers"
	"github.com/Kunniii/go_gin_gorm_test/internal"
	"github.com/Kunniii/go_gin_gorm_test/middlewares"
	"github.com/gin-gonic/gin"
)

// this function will run before main()
func init() {
	gin.SetMode(gin.ReleaseMode)
	internal.LoadEnv()
	internal.ConnectDB()
	internal.AutoMigrate()
}

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})

	postsRouter := router.Group("/posts")
	{
		postsRouter.POST("/", controllers.CreatePost)
		postsRouter.GET("/", controllers.GetAllPosts)

		postsRouter.GET("/:id", controllers.GetPostById)
		postsRouter.PUT("/:id", controllers.UpdatePost)
		postsRouter.DELETE("/:id", controllers.DeletePostById)
	}

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
		userRouter.GET("/validate", middlewares.CheckAuth, controllers.Login)
	}

	log.Fatal(router.Run())
}
