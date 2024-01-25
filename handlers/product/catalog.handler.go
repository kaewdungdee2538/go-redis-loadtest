package handlers

import (
	services "go-redis/services/product"

	"github.com/gofiber/fiber/v2"
)

type catalogHandler struct {
	catalogSrv services.CatalogService
}

func NewCatalogHandler(catalogSrv services.CatalogService) CatalogHandler {
	return catalogHandler{catalogSrv}
}

func (h catalogHandler) GetProducts(c *fiber.Ctx) error {

	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status": "ok",
		"data":   products,
	}

	return c.JSON(response)
}
