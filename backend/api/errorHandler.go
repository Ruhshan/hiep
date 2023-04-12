package api

import (
	"encoding/json"
	"fmt"
	"github.com/Ruhshan/hiep/backend/pkg/errorMessages"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}


func ErrorHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Next()

	var allErrors = c.Errors
	for _, err := range allErrors {
		// log, handle, etc.
		fmt.Println(err)
	}

	var lastError = allErrors.Last()

	var status = http.StatusInternalServerError
	var apiError ErrorResponse


	if lastError!=nil{
		switch lastError.Err.(type) {
			case *json.UnmarshalTypeError:
				status = http.StatusBadRequest
				apiError.Error = "Invalid value at: "+lastError.Err.(*json.UnmarshalTypeError).Field
				break
			case error:
				apiError = getMessageForError(lastError.Err)
				break
			default:
				apiError.Error="Unknown Error"
			}
		c.JSON(status, apiError)
	}

}

func getMessageForError(e error)  ErrorResponse{
	switch e {
	case errorMessages.ErrContainsInvalidCharacters:
		return ErrorResponse{"Input Contains invalid character"}
	default:
		return ErrorResponse{"Unknown Error"}
	}


}
