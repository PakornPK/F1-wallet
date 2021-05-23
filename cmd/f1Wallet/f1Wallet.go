package main

import (
	"github.com/PakornPK/F1-wallet/internal/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	route.Setup(app)
	app.Listen(":3000")
}
