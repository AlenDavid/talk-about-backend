package plants

import (
	"github.com/gin-gonic/gin"
)

type actions interface {
}

type services struct {
	actions
}

func Group(action actions, base *gin.RouterGroup) {
	c := services{action}

	plants := base.Group("plants")

	plants.GET("", c.Latest)
	plants.POST("", c.Create)
	plants.PUT(":id", c.Update)
}
