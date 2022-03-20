package validator

import (
	"ack/go-hello-web/app/models"
	"ack/go-hello-web/framework"
	"fmt"
)

func CheckOneUser(account string) bool {
	result := false
	var user models.User

	dbResult := framework.SqlSession.Where("account = ?", account).Find(&user)
	if dbResult.Error != nil {
		fmt.Printf("Get User Info Failed:%v\n", dbResult.Error)
	} else {
		result = true
	}
	return result
}
