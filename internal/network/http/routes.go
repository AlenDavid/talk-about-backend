package http

import (
	"github.com/alendavid/talk-about-backend/internal/network/http/plants"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Network
	Base, Version string
}

func (routes Routes) Default(engine *gin.Engine) error {
	base := engine.Group(routes.Base).Group(routes.Version)

	plants.Group(routes.database, base)

	return nil
}
