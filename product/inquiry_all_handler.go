package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type inquiryAllHandler struct {
	DB *gorm.DB
}

func NewInquiryAllHandler(DB *gorm.DB) *inquiryAllHandler {
	return &inquiryAllHandler{DB: DB}
}

func (handler *inquiryAllHandler) InquiryAll(c *fiber.Ctx) error {
	var entList []Product
	result := handler.DB.Find(&entList)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(&entList)
}
