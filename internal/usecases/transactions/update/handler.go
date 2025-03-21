package update

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiErrors "github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	apiContext "github.com/melegattip/financial-resume-engine/internal/infrastructure/context"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
)

const updateTransaction = "updateTransaction"

type UpdateTransaction struct {
	Service Update
}

func (handler *UpdateTransaction) Handle(c *gin.Context) error {
	ctx := apiContext.SetAction(c.Request.Context(), updateTransaction)

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

	var updateRequest UpdateTransactionRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		logger.Error(ctx, err, err.Error(), logs.Tags{})
		return apiErrors.NewBadRequest(logs.ErrorBinding.GetMessageWithMapParams(logs.Params{"error": err}))
	}

	serviceResponse, err := handler.Service.Execute(ctx, userID, transactionID, &updateRequest)
	if err != nil {
		logger.Error(ctx, err, logs.ErrorUpdatingTransaction.GetMessage(), logs.Tags{"error": err})
		return err
	}

	c.JSON(http.StatusOK, serviceResponse)
	return nil
}
