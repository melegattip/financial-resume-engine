package list

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiErrors "github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	apiContext "github.com/melegattip/financial-resume-engine/internal/infrastructure/context"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
)

const listTransactions = "listTransactions"

type ListTransactions struct {
	Service List
}

func (handler *ListTransactions) Handle(c *gin.Context) error {
	ctx := apiContext.SetAction(c.Request.Context(), listTransactions)

	userID := c.GetHeader("x-caller-id")
	if userID == "" {
		logger.Error(ctx, nil, logs.ErrorUnauthorizedRequest.GetMessageWithMapParams(
			logs.Params{"x-caller-id": userID}),
			logs.Tags{})
		return apiErrors.NewUnauthorizedRequest(logs.ErrorUnauthorizedRequest.GetMessage())
	}

	serviceResponse, err := handler.Service.Execute(ctx, userID)
	if err != nil {
		logger.Error(ctx, err, logs.ErrorListingTransactions.GetMessage(), logs.Tags{"error": err})
		return err
	}

	c.JSON(http.StatusOK, serviceResponse)
	return nil
}
