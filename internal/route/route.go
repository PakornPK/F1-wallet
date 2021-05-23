package route

import (
	"github.com/PakornPK/F1-wallet/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/wallet", controller.GetWalletInfo)
}
