package main

import (
	"kington/adapters"
	"kington/core"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to conncet database")
	}
	db.AutoMigrate(&core.Order{})

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	app.Listen(":8000")
}
