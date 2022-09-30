package routes

import (
	"context"
	"fmt"
	"golang-ecommerce-server/models"
	"golang-ecommerce-server/responses"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email of password is incorrect")
		check = false
	}
	return check, msg
}

func AuthRoutes(client *mongo.Client, router *gin.Engine) {
	var userCollection = client.Database("go-ecommerce").Collection("users")
	var validate = validator.New()

	router.POST("/api/signup", func(c *gin.Context) {
		log.Println("==== SIGNUP ====")
		var user models.User
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

		// validate the request body
		if err := c.BindJSON(&user); err != nil {
			log.Fatal(err.Error())
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			log.Fatal(validationErr.Error())
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		// defer cancel()

		if err != nil {
			log.Fatal(err.Error())
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "phone number already exists", Data: map[string]interface{}{"error": "this phoen number is already exists"}})
			return
		}

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		// token, refreshToken, _ := hel
	})

}
