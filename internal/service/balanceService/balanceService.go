package balanceService

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Zauberstuhl/go-coinbase"
)

func Get() (string, error) {
	cb := coinbase.APIClient{
		Key:    "9Ht7B5Ut5U7Is1T1",
		Secret: "roHxAUbanXQGxNu7EZHheE1axL8VnFib",
	}

	acc, err := cb.Accounts()
	if err != nil {
		return "", errors.New("Get : Error Account")
	}
	for _, acc := range acc.Data {
		if strings.EqualFold(acc.Currency, "BTC") {
			addr, _ := cb.GetAddresses(acc.Id)
			var address string
			for _, addr := range addr.Data {
				address = addr.Address
			}
			return fmt.Sprintf("Address: %s\nName: %s\nType: %s\nAmount: %f\nCurrency: %s\n",
				address, acc.Name, acc.Type,
				acc.Balance.Amount, acc.Balance.Currency), nil
		}
	}
	return "", nil
}
