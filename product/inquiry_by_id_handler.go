package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type inquiryByIdHandler struct {
	DB *gorm.DB
}

func NewInquiryByIdHandler(DB *gorm.DB) *inquiryByIdHandler {
	return &inquiryByIdHandler{DB: DB}
}

func (handler *inquiryByIdHandler) InquiryById(c *fiber.Ctx) error {
	id := c.Params("id")
	var ent Product
	result := handler.DB.First(&ent, id)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(&ent)
}