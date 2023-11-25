package http

import (
	"fmt"
	"net/http"

	"github.com/alendavid/talk-about-backend/internal/broker"
	"github.com/gin-gonic/gin"
)

type Database interface {
}

type Network struct {
	database Database
	broker   broker.Broker
}

func New(db Database, broker broker.Broker) Network {
	return Network{db, broker}
}

func (n *Network) Run() error {
	routes := Routes{*n, "/api", "v1"}

	gin.SetMode(gin.DebugMode)

	engine := gin.Default()
	engine.ForwardedByClientIP = true
	engine.SetTrustedProxies([]string{"0.0.0.0"})

	routes.Default(engine)

	if gin.IsDebugging() {
		engine.NoRoute(func(c *gin.Context) {
			c.Status(http.StatusOK)

			info := engine.Routes()
			for _, route := range info {
				fmt.Fprintf(c.Writer, "%-6s %-25s --> %s\n", route.Method, route.Path, route.Handler)
			}
		})
	}

	return engine.Run("0.0.0.0:8080")
}
