package routes

import (
	"context"
	"fmt"
	"golang-ecommerce-server/models"
	"golang-ecommerce-server/responses"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProductRoutes(client *mongo.Client, router *gin.Engine) {
	var productCollection = client.Database("go-ecommerce").Collection("products")

	// ==========================
	// ==== GET ALL PRODUCTS ====
	// ==========================
	router.GET("/get/all/products", func(c *gin.Context) {
		var products []models.Product
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		cursor, err := productCollection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "Not Found"})
			return
		}
		// for cursor.Next(ctx) {
		// 	var result bson.D
		// 	if err := cursor.Decode(&result); err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(result)
		// }

		for cursor.Next(ctx) {
			var element models.Product
			err := cursor.Decode(&element)
			if err != nil {
				fmt.Println(err)
			}
			products = append(products, element)
		}
		log.Print(products)
		cursor.Close(ctx)
		c.JSON(http.StatusOK, products)
	})

	// ===========================================
	// ==== GET PRODUCTS USING LIMIT AND SKIP ====
	// ===========================================
	router.GET("/get/all/products/:limit/:skip", func(c *gin.Context) {
		limit, err := strconv.ParseInt(c.Param("limit"), 0, 8)
		skip, err := strconv.ParseInt(c.Param("skip"), 0, 8)
		options := options.Find()
		options.SetSkip(skip)
		options.SetLimit(limit)
		var products []models.Product
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		cursor, err := productCollection.Find(ctx, bson.D{}, options)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "Not Found"})
			return
		}
		for cursor.Next(ctx) {
			var element models.Product
			err := cursor.Decode(&element)
			if err != nil {
				fmt.Println(err)
			}
			products = append(products, element)
		}
		log.Print(products)
		cursor.Close(ctx)
		c.JSON(http.StatusOK, products)
	})

	// ========================
	// ==== SEARCH PRODUCT ====
	// ========================
	router.GET("/get/search/products/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Print(name)
		var products []models.Product
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		cursor, err := productCollection.Find(ctx, bson.M{"title": bson.M{"$regex": name, "$options": "i"}})
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "Not Found"})
			return
		}
		for cursor.Next(ctx) {
			var element models.Product
			err := cursor.Decode(&element)
			if err != nil {
				fmt.Println(err)
			}
			products = append(products, element)
		}
		cursor.Close(ctx)
		c.JSON(http.StatusOK, products)
	})

	// ===========================
	// ==== GET PRODUCT BY ID ====
	// ===========================
	router.GET("/get/product/:id", func(c *gin.Context) {
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		fmt.Print(id)
		var product models.Product
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		err := productCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "Not Found"})
			return
		}
		c.JSON(http.StatusOK, product)
	})
}
