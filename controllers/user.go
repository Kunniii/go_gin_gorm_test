package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Kunniii/go_gin_gorm_test/internal"
	"github.com/Kunniii/go_gin_gorm_test/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

func Register(context *gin.Context) {
	var reqBody struct {
		Email    string
		Password string
	}

	if context.Bind(&reqBody) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Make sure to put the JSON key as String",
		})
		return
	}

	fmt.Print(reqBody.Email, reqBody.Password)

	// hash user's password
	hashByte, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Password hashing error!",
		})
		return
	}

	user := models.User{Email: reqBody.Email, Password: string(hashByte)}

	if result := internal.DB.Create(&user); result.Error != nil {
		err := result.Error
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": "Email already exists!",
			})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Cannot create user!",
			})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

func Login(context *gin.Context) {
	var reqBody struct {
		Email    string
		Password string
	}

	if err := context.Bind(&reqBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Make sure to put JSON key as String!",
		})
		return
	}

	var user models.User
	result := internal.DB.First(&user, "email = ?", reqBody.Email)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid credential!",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid credential!",
		})
		return
	}

	// token expires in 30 days
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"expires": time.Now().Add(time.Hour * 24 * 30).Unix(),
		"email":   user.Email,
	})

	if tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET"))); err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Unable to create Token!",
		})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{
			"msg":   "OK",
			"token": tokenString,
			"email": user.Email,
			"id":    user.ID,
		})
	}
}

func Validate(context *gin.Context) {

	uid, _ := context.Get("userID")
	email, _ := context.Get("userEmail")

	context.JSON(http.StatusOK, gin.H{
		"msg": "OK",
		"data": map[string]any{
			"id":    uid,
			"email": email,
		},
	})
}
