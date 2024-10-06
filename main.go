package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"building-system/config"
	"building-system/database"
	"building-system/routes"
)

func main() {
	config := config.Load()

	db, err := database.Init(config.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()

	routes.RegisterRoutes(app, db)

	log.Fatal(app.Listen(config.ServerPort))
}
