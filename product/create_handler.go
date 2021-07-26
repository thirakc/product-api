package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type createHandler struct {
	DB *gorm.DB
}

func NewCreateHandler(DB *gorm.DB) *createHandler {
	return &createHandler{DB: DB}
}

type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   *bool   `json:"available"`
}

func (handler *createHandler) Create(c *fiber.Ctx) error {
	var req CreateRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(nil)
	}

	ent := Product{
		Name:        req.Name,
		Description: req.Description,
		Available:   req.Available,
	}
	_ = handler.DB.Create(&ent)

	return c.Status(http.StatusOK).JSON(nil)
}
