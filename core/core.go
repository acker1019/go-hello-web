package core

import (
	"ack/go-hello-web/core/settings"
	"ack/go-hello-web/core/urls"

	"github.com/gin-gonic/gin"
)

type Core struct{}

func (core *Core) GetDbConfig() string {
	return settings.DbConfig
}

func (core *Core) SetUrl(router *gin.Engine) {
	urls.V1(router)
}
