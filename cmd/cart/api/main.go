package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/scastillop/fiber_skeleton/internal/api/routes"
)

func init() {
}

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	var port string
	app := fiber.New(fiber.Config{
		ReadBufferSize: 9000,
	})

	routes.SetupRoutes(app)

	if port = os.Getenv("HTTP_PORT"); port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
