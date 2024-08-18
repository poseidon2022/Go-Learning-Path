package main

import (
	route "blog_filtration_feature/routes"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
	router := gin.Default()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("error while connecting to mongo")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("error while ensuring connection")
	}
	
	dataBase := client.Database("blog")
	fmt.Println("Database connected")
	route.Setup(router, dataBase)
	router.Run("localhost:8080")
}