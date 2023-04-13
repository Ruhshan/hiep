package api

import (
	"github.com/Ruhshan/hiep/backend/config"
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes() routes {
	gin.SetMode(config.Get("ginMode"))
	r := routes{
		router: gin.Default(),
	}
	r.router.Use(ErrorHandler)

	return r
}

func (r routes) GetBaseRoute(path string) *gin.RouterGroup{
	return r.router.Group("v1/"+path)
}


func (r routes) Run(addr ...string) error {
	port := config.Get("port")
	return r.router.Run(":"+port)
}

