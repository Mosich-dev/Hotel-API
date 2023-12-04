package main

import (
	"context"
	"flag"
	"github.com/Mosich-dev/Hotel-API/API"
	"github.com/Mosich-dev/Hotel-API/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	dburi = "mongodb://localhost:27017"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{
			"error": err.Error(),
		})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(config)
	listenAddr := flag.String("listenaddr", ":5000", "The Listen address of the API server")
	flag.Parse()

	apiv1 := app.Group("/api/v1")

	// Handler init
	userHandler := API.NewUserHandler(db.NewMongoUserStore(client))

	// Routing
	apiv1.Get("/foo", handleFoo)
	app.Get("/eee", userHandler.HandleGetUsers)
	app.Get("/eee/:id", userHandler.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleFoo(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{"date": "6th azar", "id": "0"})

}
