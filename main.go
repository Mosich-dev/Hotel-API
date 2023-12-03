package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	listenAddr := flag.String("listenaddr", ":5000", "The Listen address of the API server")
	flag.Parse()

	apiv1 := app.Group("/api/v1")

	apiv1.Get("/foo", handleFoo)
	app.Get("/eee", handleEee)
	app.Listen(*listenAddr)
}

func handleEee(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{"name": "Mosi", "id": "44"})

}

func handleFoo(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{"date": "6th azar", "id": "0"})

}
