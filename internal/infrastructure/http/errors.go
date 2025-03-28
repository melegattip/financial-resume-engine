package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melegattip/financial-resume-engine/internal/core/errors"
)

// ErrorResponse representa la estructura de respuesta de error
type ErrorResponse struct {
	Error   string `json:"error" example:"Error message"`
	Message string `json:"message" example:"Bad Request"`
}

func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *errors.BadRequest:
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   e.Error(),
			Message: "Bad Request",
		})
	case *errors.UnauthorizedRequest:
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   e.Error(),
			Message: "Unauthorized",
		})
	case *errors.ResourceNotFound:
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   e.Error(),
			Message: "Not Found",
		})
	case *errors.TooManyRequests:
		c.JSON(http.StatusTooManyRequests, ErrorResponse{
			Error:   e.Error(),
			Message: "Too Many Requests",
		})
	case *errors.ResourceAlreadyExists:
		c.JSON(http.StatusConflict, ErrorResponse{
			Error:   e.Error(),
			Message: "Resource Already Exists",
		})
	default:
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "Internal Server Error",
		})
	}
}
