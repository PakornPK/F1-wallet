package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PakornPK/F1-wallet/internal/models"
	"github.com/Zauberstuhl/go-coinbase"
)

func GetWallet(userAccount models.UserAccount) ([]models.Wallet, error) {
	var wallet []models.Wallet
	coinbase := coinbase.APIClient{
		Key:    *userAccount.AccessKey,
		Secret: *userAccount.SecretKey,
	}

	coinbaseAccounts, err := coinbase.Accounts()
	if len(coinbaseAccounts.Data) == 0 {
		errorMessage := fmt.Sprintf(models.ACCOUNT_NOT_AVAILABLE, err)
		return nil, errors.New(errorMessage)
	}
	for _, coinbaseAccount := range coinbaseAccounts.Data {
		if strings.EqualFold(coinbaseAccount.Currency, "BTC") {
			coinbaseAddresses, _ := coinbase.GetAddresses(coinbaseAccount.Id)
			coinbaseAddress := coinbaseAddresses.Data[0].Address

			wallet = append(wallet, models.Wallet{
				Address:  &coinbaseAddress,
				Name:     &coinbaseAccount.Name,
				Type:     &coinbaseAccount.Type,
				Balance:  &coinbaseAccount.Balance.Amount,
				Currency: &coinbaseAccount.Balance.Currency,
			})
		}
	}
	return wallet, nil
}
