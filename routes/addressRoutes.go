package routes

import (
	"context"
	"fmt"
	"golang-ecommerce-server/middleware"
	"golang-ecommerce-server/models"
	"golang-ecommerce-server/responses"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddressRoutes(client *mongo.Client, router *gin.Engine) {
	var addressCollection = client.Database("go-ecommerce").Collection("addresses")

	// ==============================
	// ==== GET MY ALL ADDRESSES ====
	// ==============================
	router.GET("/get/all/addresses", middleware.Authenticate(), func(c *gin.Context) {
		var addresses []models.Address
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		cursor, err := addressCollection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "Not Found"})
			return
		}
		for cursor.Next(ctx) {
			var element models.Address
			err := cursor.Decode((&element))
			if err != nil {
				fmt.Println(err)
			}
			addresses = append(addresses, element)
		}
		log.Print(addresses)
		//close teh cursor one finished
		cursor.Close(ctx)
		c.JSON(http.StatusOK, addresses)
	})

	// =====================
	// ==== ADD ADDRESS ====
	// =====================
	router.POST("/add/address", middleware.Authenticate(), func(c *gin.Context) {
		log.Println("==== ADD ADDRESS ====")
		fmt.Println(c.GetString("id"))
		userId, err := primitive.ObjectIDFromHex(c.GetString("id"))
		log.Print(userId)
		var address models.Address
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

		// validate the request body
		if err := c.BindJSON(&address); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newAddress := models.Address{
			ID:      primitive.NewObjectID(),
			UserId:  userId,
			Name:    address.Name,
			Address: address.Address,
			City:    address.City,
			State:   address.State,
			Email:   address.Email,
			Phone:   address.Phone,
		}

		fmt.Println(newAddress)
		result, err := addressCollection.InsertOne(ctx, newAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	})
}
