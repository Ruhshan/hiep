package api

import (
	"github.com/Ruhshan/hiep/backend/config"
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewRoutes() routes {
	gin.SetMode(config.Get("ginMode"))
	r := routes{
		router: gin.Default(),
	}
	r.router.Use(ErrorHandler)
	r.router.Use(CORSMiddleware())

	return r
}

func (r routes) GetBaseRoute(path string) *gin.RouterGroup{
	return r.router.Group("/api/v1/"+path)
}


func (r routes) Run(addr ...string) error {
	port := config.Get("port")
	return r.router.Run(":"+port)
}

