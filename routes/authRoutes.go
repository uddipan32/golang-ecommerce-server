package routes

import (
	"context"
	"fmt"
	helper "golang-ecommerce-server/helpers"
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
		fmt.Println(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("phone number or password is incorrect")
		check = false
	}
	return check, msg
}

func AuthRoutes(client *mongo.Client, router *gin.Engine) {
	var userCollection = client.Database("go-ecommerce").Collection("users")
	var validate = validator.New()

	// ================
	// ==== SIGNUP ====
	// ================
	router.POST("/api/signup", func(c *gin.Context) {
		log.Println("==== SIGNUP ====")
		var user models.User
		ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

		// validate the request body
		if err := c.BindJSON(&user); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			fmt.Println(validationErr.Error())
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		// defer cancel()

		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "phone number already exists", Data: map[string]interface{}{"error": "this phoen number is already exists"}})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		var User_id = user.ID.Hex()
		token, _ := helper.GenerateAllTokens(*user.Email, *user.Name, *user.Phone, User_id)

		_, insertErr := userCollection.InsertOne(ctx, user) // RETURNS INSERTED ID
		if insertErr != nil {
			fmt.Println(insertErr.Error())
			msg := fmt.Sprintf("User itm was not created")

			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: msg, Data: map[string]interface{}{"error": msg}})
		}
		// defer cancel()
		var result = map[string]interface{}{
			"_id":   user.ID,
			"name":  *user.Name,
			"phone": *user.Phone,
			"email": *user.Email,
			"token": token,
		}
		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	})
	// ================
	// ==== SIGNIN ====
	// ================
	router.POST("/api/signin", func(c *gin.Context) {
		log.Println("==== SIGNIN ====")
		var user models.User
		var foundUser models.User
		ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

		// validate the request body
		if err := c.BindJSON(&user); err != nil {
			print("INVALID JSON FORMAT")
			print(err.Error())
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"phone": user.Phone}).Decode(&foundUser)
		// defer cancel()
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Data: map[string]interface{}{"data": map[string]interface{}{"error": "phone number or password is incorrect"}}})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		if passwordIsValid != true {
			fmt.Println(msg)
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Data: map[string]interface{}{"data": map[string]interface{}{"error": "phone number or password is incorrect"}}})
			return
		}

		err = userCollection.FindOne(ctx, bson.M{"_id": foundUser.ID}).Decode(&foundUser)
		if foundUser.Phone == nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Data: map[string]interface{}{"data": map[string]interface{}{"error": "user not found"}}})
		}
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Data: map[string]interface{}{"data": map[string]interface{}{"error": "phone number or password is incorrect"}}})
			return
		}
		var User_id = foundUser.ID.Hex()
		token, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.Name, *foundUser.Phone, User_id)

		var result = map[string]interface{}{
			"_id":   foundUser.ID,
			"name":  foundUser.Name,
			"phone": foundUser.Phone,
			"email": foundUser.Email,
			"token": token,
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	})
}
