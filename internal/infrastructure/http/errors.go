package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melegattip/financial-resume-engine/internal/core/errors"
)

// ErrorResponse representa la estructura de respuesta de error
type ErrorResponse struct {
	Error   string `json:"error" example:"Error message"`
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *errors.BadRequest:
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   e.Error(),
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
		})
	case *errors.UnauthorizedRequest:
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   e.Error(),
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	case *errors.ResourceNotFound:
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   e.Error(),
			Code:    http.StatusNotFound,
			Message: "Not Found",
		})
	case *errors.TooManyRequests:
		c.JSON(http.StatusTooManyRequests, ErrorResponse{
			Error:   e.Error(),
			Code:    http.StatusTooManyRequests,
			Message: "Too Many Requests",
		})
	default:
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Internal server error",
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}
}
