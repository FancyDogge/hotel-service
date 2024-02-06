package main

import (
	"flag"

	"github.com/FancyDogge/hotel-service/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "The listen port of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUser)

	app.Listen(*listenAddr)
}
