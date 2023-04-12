package api

import (
	"errors"
	"github.com/Ruhshan/hiep/backend/models/requests"
	"github.com/Ruhshan/hiep/backend/pkg/errorMessages"
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
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var res, err = i.processor.ProcessPayload(request)
	if err != nil {
		handleErrors(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func handleErrors(c *gin.Context, err error) {
	if errors.Is(err, errorMessages.ErrContainsInvalidCharacters) {
		c.AbortWithError(http.StatusBadRequest, err)
	} else if errors.Is(err, errorMessages.ErrContainsMoreThanOneFastaSequence) {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}