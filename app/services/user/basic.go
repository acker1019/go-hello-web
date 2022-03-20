package user

import (
	"ack/go-hello-web/app/models"
	"ack/go-hello-web/app/validator"
	"ack/go-hello-web/framework"
	"fmt"
)

func RegisterOneUser(account string, password string, email string) error {
	if !validator.CheckOneUser(account) {
		return fmt.Errorf("User exists.")
	}
	user := models.User{
		Account:  account,
		Password: password,
		Email:    email,
	}
	insertErr := framework.SqlSession.Model(&models.User{}).Create(&user).Error
	return insertErr
}
