package api

import (
	"github.com/Ruhshan/hiep/backend/models/requests"
	"github.com/Ruhshan/hiep/backend/pkg/service/instanthiep"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InstantHiepRoutes struct {
	rg *gin.RouterGroup
	processor instanthiep.HiepProcessor
}

func ConfigureInstantHiepRoutes(rg *gin.RouterGroup, processor instanthiep.HiepProcessor)  {
	hiepRoutes := InstantHiepRoutes{rg, processor}
	hiepRoutes.configure()

}

func (i InstantHiepRoutes) configure()  {
	i.rg.POST("calculate", i.calculate)

}

func (i InstantHiepRoutes) calculate(c *gin.Context)  {
	var request requests.InstantHiepRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res, err = i.processor.ProcessPayload(request);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}