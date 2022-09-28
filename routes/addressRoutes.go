package routes

import (
	"context"
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

	router.GET("/get/all/addresses", func(c *gin.Context) {
		var addresses []models.Address
		collection := client.Database("go-ecommerce").Collection("addresses")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		cursor, err := collection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "No Found"})
			return
		}
		for cursor.Next(ctx) {
			var element models.Address
			err := cursor.Decode((&element))
			if err != nil {
				log.Fatal(err)
			}
			addresses = append(addresses, element)
		}
		log.Print(addresses)
		//close teh cursor one finished
		cursor.Close(ctx)
		c.JSON(http.StatusOK, addresses)
	})

	router.POST("/add/address", func(c *gin.Context) {
		log.Println("==== ADD ADDRESS ====")
		var address models.Address
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

		// validate the request body
		if err := c.BindJSON(&address); err != nil {
			log.Fatal(err.Error())
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		log.Print(address)

		newAddress := models.Address{
			ID:      primitive.NewObjectID(),
			Name:    address.Name,
			Address: address.Address,
			City:    address.City,
			State:   address.State,
			Email:   address.Email,
			Phone:   address.Phone,
		}

		log.Print(newAddress)

		collection := client.Database("go-ecommerce").Collection("addresses")
		result, err := collection.InsertOne(ctx, newAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	})
}
