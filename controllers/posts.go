package controllers

import (
	"github.com/Kunniii/go_gin_gorm_test/internal"
	"github.com/Kunniii/go_gin_gorm_test/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(context *gin.Context) {
	// get request body

	var reqBody models.Post

	if err := context.Bind(&reqBody); err != nil {
		context.JSON(400, gin.H{
			"msg": "Make sure to put JSON key as String!",
		})
		return
	}

	// create a new post
	post := models.Post{Title: reqBody.Body, Body: reqBody.Body}

	// save to database
	result := internal.DB.Create(&post)

	// handle if error
	if result.Error != nil {
		context.JSON(400, gin.H{
			"msg": "Could not create Post!",
		})
		return
	}

	// return a JSON of that post

	context.JSON(200, gin.H{
		"msg":  "OK",
		"data": post,
	})
}

func GetAllPosts(context *gin.Context) {
	var posts []models.Post
	internal.DB.Find(&posts)

	context.JSON(200, gin.H{
		"msg":  "OK",
		"data": posts,
	})

}

func GetPostById(context *gin.Context) {
	// Get id from URL
	id := context.Param("id")

	// find it in database
	var post models.Post
	if result := internal.DB.First(&post, id); result.Error != nil {
		context.JSON(200, gin.H{
			"msg":  "OK",
			"data": struct{}{},
		})
	} else {
		context.JSON(200, gin.H{
			"msg":  "OK",
			"data": post,
		})
	}
}

func UpdatePost(context *gin.Context) {
	// get id from URL
	id := context.Param("id")

	// get update value
	var reqBody models.Post

	if err := context.Bind(&reqBody); err != nil {
		context.JSON(400, gin.H{
			"msg": "Make sure to put JSON key as String!",
		})
		return
	}

	// find it in database
	var post models.Post
	if result := internal.DB.First(&post, id); result.Error != nil {
		// error
		context.JSON(400, gin.H{
			"msg": "ID not found",
		})
		return
	}

	internal.DB.Model(&post).Updates(models.Post{
		Title: reqBody.Title,
		Body:  reqBody.Body,
	})

	context.JSON(200, gin.H{
		"msg":  "OK",
		"data": post,
	})
}

func DeletePostById(context *gin.Context) {
	// get id from URL
	id := context.Param("id")

	if result := internal.DB.Delete(&models.Post{}, id); result.Error != nil {
		context.JSON(400, gin.H{
			"msg": result.Error,
		})
	} else {
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	}
}
