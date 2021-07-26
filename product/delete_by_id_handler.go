package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type deleteByIdHandler struct {
	DB *gorm.DB
}

func NewDeleteByIdHandler(DB *gorm.DB) *deleteByIdHandler {
	return &deleteByIdHandler{DB: DB}
}

func (handler *deleteByIdHandler) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	result := handler.DB.Delete(Product{}, id)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(nil)
}
