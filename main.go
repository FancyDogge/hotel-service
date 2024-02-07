package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/FancyDogge/hotel-service/api"
	"github.com/FancyDogge/hotel-service/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://admin:adminpassword@localhost:27017/"
const dbName = "hotel-service"
const userCollection = "users"

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "The listen port of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	coll := client.Database(dbName).Collection(userCollection)

	user := types.User{
		FirstName: "Pepe",
		LastName:  "The Frog",
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	var james types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&james); err != nil {
		log.Fatal(err)
	}
	fmt.Println(james)
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUser)
	apiv1.Get("/user/:id", api.HandleGetUsers)

	app.Listen(*listenAddr)
}
