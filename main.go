package main

import (
	"context"
	"golang-ecommerce-server/routes"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	port := os.Getenv("PORT")
	// router := gin.New() 	// WITH OUT LOGS
	router := gin.Default() // WITH DEFAULT LOGS
	client := getClient()

	// ==== ROUTES ====
	routes.AuthRoutes(client, router)
	routes.AddressRoutes(client, router)
	routes.ProductRoutes(client, router)
	router.Run("localhost:" + port)
}

func getClient() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	println("CONNECTED")
	print(client)
	return client
}
