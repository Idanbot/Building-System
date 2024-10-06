package apartment

import (
	"building-system/database"
	"building-system/models"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// GetApartments retrieves all apartments
func GetApartments(c *fiber.Ctx) error {
	apartments, err := models.Apartments().All(c.Context(), database.DB)
	if err != nil {
		return c.Status(500).SendString("Error retrieving apartments")
	}
	return c.JSON(apartments)
}

// GetApartment retrieves a single apartment by ID
func GetApartment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	apartment, err := models.FindApartment(c.Context(), database.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).SendString("Apartment not found")
		}
		return c.Status(500).SendString("Error retrieving apartment")
	}
	return c.JSON(apartment)
}

// GetApartmentsByBuilding retrieves apartments for a specific building
func GetApartmentsByBuilding(c *fiber.Ctx) error {
	buildingID, err := strconv.Atoi(c.Params("building_id"))
	if err != nil {
		return c.Status(400).SendString("Invalid building ID format: " + err.Error() + " received: " + c.Params("building_id"))
	}

	apartments, err := models.Apartments(models.ApartmentWhere.BuildingID.EQ(buildingID)).All(c.Context(), database.DB)
	if err != nil {
		return c.Status(500).SendString("Error retrieving apartments")
	}
	if len(apartments) == 0 {
		return c.Status(404).SendString("No apartments found for the given building")
	}
	return c.JSON(apartments)
}

// CreateApartment creates a new apartment
func CreateApartment(c *fiber.Ctx) error {
	var apartment models.Apartment
	if err := c.BodyParser(&apartment); err != nil {
		return c.Status(400).SendString("Invalid request body: " + err.Error())
	}

	if err := apartment.Insert(c.Context(), database.DB, boil.Infer()); err != nil {
		return c.Status(500).SendString("Error creating apartment: " + err.Error())
	}
	return c.Status(201).JSON(apartment)
}

// DeleteApartment deletes an apartment by ID
func DeleteApartment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	// Find the apartment by ID
	apartment, err := models.FindApartment(c.Context(), database.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).SendString("Apartment not found")
		}
		return c.Status(500).SendString("Error retrieving apartment")
	}

	// Delete the apartment
	rowsAff, err := apartment.Delete(c.Context(), database.DB)
	if err != nil {
		return c.Status(500).SendString("Error deleting apartment")
	}

	if rowsAff == 0 {
		return c.Status(404).SendString("Apartment not found")
	}

	return c.SendStatus(204)
}
