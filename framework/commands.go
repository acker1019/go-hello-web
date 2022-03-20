package framework

import (
	"ack/go-hello-web/app/models"

	"github.com/gin-gonic/gin"
)

func RunServer(core Core, addr string) {
	// launch db driver
	db, ormErr := Initialize(core.GetDbConfig())
	if ormErr != nil {
		panic(ormErr)
	}
	migrateErr := db.AutoMigrate(&models.User{})
	if migrateErr != nil {
		return
	}

	// launch gin-gonic app
	router := gin.Default()
	core.SetUrl(router)
	router.Run(addr)
}
