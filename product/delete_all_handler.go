package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type deleteAllHandler struct {
	DB *gorm.DB
}

func NewDeleteAllHandler(DB *gorm.DB) *deleteAllHandler {
	return &deleteAllHandler{DB: DB}
}

func (handler *deleteAllHandler) DeleteAll(c *fiber.Ctx) error {

}
