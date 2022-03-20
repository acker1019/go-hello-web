package framework

import "github.com/gin-gonic/gin"

type Core interface {
	// getter
	GetDbConfig() string

	// setter
	SetUrl(router *gin.Engine)
}
