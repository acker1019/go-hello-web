package urls

import (
	app_views "ack/go-hello-web/app/views"

	"github.com/gin-gonic/gin"
)

func groupUser(group *gin.RouterGroup) *gin.RouterGroup {
	user := group.Group("/user")
	{
		user.POST("/login", app_views.Login)
	}
	return user
}

func V1(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/calculate/:action", app_views.Calculate)
		groupUser(v1)
	}
}
