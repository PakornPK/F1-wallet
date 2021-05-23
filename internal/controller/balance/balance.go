package balance

import (
	"errors"

	"github.com/PakornPK/F1-wallet/internal/service/balanceService"
	"github.com/gofiber/fiber/v2"
)

func GetBalance(c *fiber.Ctx) error {
	bl, err := balanceService.Get()
	if err != nil {
		return errors.New("GetBalace : " + err.Error())
	}
	return c.SendString(bl)
}
