package api

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes() routes {
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
	return r.router.Run()
}

