package models

type UserAccount struct {
	AccessKey *string `json:"accessKey"`
	SecretKey *string `json:"secretKey"`
}
