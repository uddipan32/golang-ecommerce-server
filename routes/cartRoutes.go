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

func CartRoutes(client *mongo.Client, router *gin.Engine) {
	var cartCollection = client.Database("go-ecommerce").Collection("carts")

	// =====================
	// ==== GET MY CART ====
	// =====================
	router.GET("/get/my/cart", middleware.Authenticate(), func(c *gin.Context) {
		log.Println("==== GET MY CART ====")
		fmt.Println(c.GetString("id"))
		var cart models.Cart
		userId, _ := primitive.ObjectIDFromHex(c.GetString(("id")))
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		err := cartCollection.FindOne(ctx, bson.M{"_user": userId}).Decode(&cart)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "Not Found"})
			return
		}
		c.JSON(http.StatusOK, cart)
	})
}
