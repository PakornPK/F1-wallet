package balance

import (
	"fmt"
	"strings"

	"github.com/Zauberstuhl/go-coinbase"
	"github.com/gofiber/fiber/v2"
)

func GetBalance(c *fiber.Ctx) error {
	cb := coinbase.APIClient{
		Key:    "9Ht7B5Ut5U7Is1T1",
		Secret: "roHxAUbanXQGxNu7EZHheE1axL8VnFib",
	}

	acc, err := cb.Accounts()
	if err != nil {
		return c.SendString("Error Account")
	}
	for _, acc := range acc.Data {
		if strings.EqualFold(acc.Currency, "BTC") {
			addr, _ := cb.GetAddresses(acc.Id)
			var address string
			for _, addr := range addr.Data {
				address = addr.Address
			}
			return c.SendString(fmt.Sprintf("Address: %s\nName: %s\nType: %s\nAmount: %f\nCurrency: %s\n",
				address, acc.Name, acc.Type,
				acc.Balance.Amount, acc.Balance.Currency))
		}
	}
	return nil

}
