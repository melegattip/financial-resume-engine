package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiErrors "github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	apiContext "github.com/melegattip/financial-resume-engine/internal/infrastructure/context"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
)

const getTransaction = "getTransaction"

type GetTransaction struct {
	Service Get
}

func (handler *GetTransaction) Handle(c *gin.Context) error {
	ctx := apiContext.SetAction(c.Request.Context(), getTransaction)

	userID := c.GetHeader("x-caller-id")
	if userID == "" {
		logger.Error(ctx, nil, logs.ErrorUnauthorizedRequest.GetMessageWithMapParams(
			logs.Params{"x-caller-id": userID}),
			logs.Tags{})
		return apiErrors.NewUnauthorizedRequest(logs.ErrorUnauthorizedRequest.GetMessage())
	}

	transactionID := c.Param("id")
	if transactionID == "" {
		logger.Error(ctx, nil, "Transaction ID is required", logs.Tags{})
		return apiErrors.NewBadRequest("Transaction ID is required")
	}

	serviceResponse, err := handler.Service.Execute(ctx, userID, transactionID)
	if err != nil {
		logger.Error(ctx, err, logs.ErrorGettingTransaction.GetMessage(), logs.Tags{"error": err})
		return err
	}

	c.JSON(http.StatusOK, serviceResponse)
	return nil
}
