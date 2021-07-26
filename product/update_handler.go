package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type updateHandler struct {
	DB *gorm.DB
}

func NewUpdateHandler(DB *gorm.DB) *updateHandler {
	return &updateHandler{DB: DB}
}

type UpdateRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   *bool   `json:"available"`
}

func (handler *updateHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	var ent Product
	result := handler.DB.First(&ent, id)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(result.Error.Error())
	}

	handler.DB.Model(&ent).Updates(Product{
		Name:        req.Name,
		Description: req.Description,
		Available:   req.Available,
	})
	return c.Status(http.StatusOK).JSON(nil)
}
