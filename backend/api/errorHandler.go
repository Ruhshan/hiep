package api

import (
	"encoding/json"
	"fmt"
	"github.com/Ruhshan/hiep/backend/pkg/errorMessages"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
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
			case validator.ValidationErrors:
				apiError = getMessageForValidationError(lastError.Err.(validator.ValidationErrors))
				break
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

func getMessageForValidationError(errs validator.ValidationErrors) ErrorResponse {

	var field = toCamelCase(errs[0].Field())
	var tag = errs[0].Tag()

	return ErrorResponse{fmt.Sprintf("Validation failed on field: %s, error: %s", field, tag)}
}

func toCamelCase(s string) string {
	if s == "" {
		return ""
	}
	lower := strings.ToLower(s)
	return strings.Replace(lower[:1]+s[1:], s[:1], lower[:1], 1)
}

func getMessageForError(e error)  ErrorResponse{

	switch e {
	case errorMessages.ErrContainsInvalidCharacters:
		return ErrorResponse{"Input contains invalid character"}
	case errorMessages.ErrContainsMoreThanOneFastaSequence:
		return ErrorResponse{"Input contains more than one fasta"}
	default:
		return ErrorResponse{"Unknown Error"}
	}


}
