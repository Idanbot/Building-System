package routes

import (
	"building-system/api/apartment"
	"building-system/api/building"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/api")

	api.Get("/buildings", building.GetBuildings)
	api.Get("/buildings/:id", building.GetBuilding)
	api.Post("/buildings", building.CreateBuilding)
	api.Delete("/buildings/:id", building.DeleteBuilding)

	api.Get("/apartments", apartment.GetApartments)
	api.Get("/apartments/:id", apartment.GetApartment)
	api.Get("/apartments/building/:buildingId", apartment.GetApartmentsByBuilding)
	api.Post("/apartments", apartment.CreateApartment)
	api.Delete("/apartments/:id", apartment.DeleteApartment)
}
