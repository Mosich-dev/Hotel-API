package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Mosich-dev/Hotel-API/API"
	"github.com/Mosich-dev/Hotel-API/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	dburi    = "mongodb://localhost:27017"
	dbname   = "hotel-api"
	userColl = "users"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client)
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)
	user := types.User{
		FirstName: "mosi",
		LastName:  "che",
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	var mosi types.User
	err = coll.FindOne(ctx, bson.M{}).Decode(&mosi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mosi)
	app := fiber.New()
	listenAddr := flag.String("listenaddr", ":5000", "The Listen address of the API server")
	flag.Parse()

	apiv1 := app.Group("/api/v1")

	apiv1.Get("/foo", handleFoo)
	app.Get("/eee", API.HandleGetUsers)
	app.Get("/eee/:id", API.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleFoo(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{"date": "6th azar", "id": "0"})

}
