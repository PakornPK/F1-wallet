package controller

import (
	"github.com/PakornPK/F1-wallet/internal/models"
	"github.com/PakornPK/F1-wallet/internal/service"
	"github.com/gofiber/fiber/v2"
)

func GetWalletInfo(c *fiber.Ctx) error {
	userAccount := new(models.UserAccount)
	if err := c.BodyParser(userAccount); err != nil {
		return err
	}

	wallet, err := service.GetWallet(*userAccount)
	if err != nil {
		// return errors.New("GetWallet: " + err.Error())
		// return fiber.NewError(fiber.StatusServiceUnavailable, "GetWallet: "+err.Error())
		return c.Status(404).JSON(models.ERROR_WALLET)
	}
	return c.Status(200).JSON(wallet)
}
