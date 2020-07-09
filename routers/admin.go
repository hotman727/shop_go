package routers

import (
	"github.com/gin-gonic/gin"
	"goshop/middlewares/roles"
)

func Admin(e *gin.Engine) {
	admin := e.Group("admin")
	admin.Use(roles.CheckAdmin())
	{
		admin.POST("/roles", roleService.CreateRole)
	}
}