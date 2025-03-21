package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melegattip/financial-resume-engine/internal/core/errors"
)

func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *errors.BadRequest:
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	case *errors.UnauthorizedRequest:
		c.JSON(http.StatusUnauthorized, gin.H{"error": e.Error()})
	case *errors.ResourceNotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
	case *errors.TooManyRequests:
		c.JSON(http.StatusTooManyRequests, gin.H{"error": e.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
