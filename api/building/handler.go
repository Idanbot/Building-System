package building

import (
	"building-system/database"
	"building-system/models"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// GetBuildings retrieves all buildings
func GetBuildings(c *fiber.Ctx) error {
	buildings, err := models.Buildings().All(c.Context(), database.DB)
	if err != nil {
		return c.Status(500).SendString("Error retrieving buildings")
	}

	return c.JSON(buildings)
}

// GetBuilding retrieves a single building by ID
func GetBuilding(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	building, err := models.FindBuilding(c.Context(), database.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).SendString("Building not found")
		}
		return c.Status(500).SendString("Error retrieving building")
	}
	return c.JSON(building)
}

// CreateBuilding creates a new building
func CreateBuilding(c *fiber.Ctx) error {
	var building models.Building
	if err := c.BodyParser(&building); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	if err := building.Insert(c.Context(), database.DB, boil.Infer()); err != nil {
		return c.Status(500).SendString("Error creating building")
	}
	return c.Status(201).JSON(building)
}

// DeleteBuilding deletes a building by ID
func DeleteBuilding(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	// Find the building by ID
	building, err := models.FindBuilding(c.Context(), database.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).SendString("Building not found")
		}
		return c.Status(500).SendString("Error retrieving building")
	}

	// Delete the building
	rowsAff, err := building.Delete(c.Context(), database.DB)
	if err != nil {
		return c.Status(500).SendString("Error deleting building")
	}

	if rowsAff == 0 {
		return c.Status(404).SendString("Building not found")
	}

	return c.SendStatus(204)
}
