package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"thirak/product-api/product"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&product.Product{})

	createHandler := product.NewCreateHandler(db)
	inquiryAllHandler := product.NewInquiryAllHandler(db)
	inquiryByIdHandler := product.NewInquiryByIdHandler(db)
	updateHandler := product.NewUpdateHandler(db)
	deleteByIdHandler := product.NewDeleteByIdHandler(db)


	api := app.Group("/api")
	api.Post("/products", createHandler.Create)
	api.Get("/products", inquiryAllHandler.InquiryAll)
	api.Get("/products/:id", inquiryByIdHandler.InquiryById)
	api.Put("/products/:id", updateHandler.Update)
	api.Delete("/products/:id", deleteByIdHandler.DeleteById)

	app.Listen(":8080")
}
