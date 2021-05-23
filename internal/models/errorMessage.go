package models

const ACCOUNT_NOT_AVAILABLE string = "The account is not available: %v"

type errorWallet struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
}

var ERROR_WALLET errorWallet = errorWallet{
	"Error",
	"Account invalid",
}