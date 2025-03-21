package delete

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiErrors "github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	apiContext "github.com/melegattip/financial-resume-engine/internal/infrastructure/context"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
)

const deleteTransaction = "deleteTransaction"

type DeleteTransaction struct {
	Service Delete
}

func NewDeleteTransactionHandler(service Delete) *DeleteTransaction {
	return &DeleteTransaction{Service: service}
}

func (h *DeleteTransaction) Handle(c *gin.Context) error {
	ctx := apiContext.SetAction(c.Request.Context(), deleteTransaction)

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

	response, err := h.Service.Execute(ctx, userID, transactionID)
	if err != nil {
		logger.Error(ctx, err, logs.ErrorDeletingTransaction.GetMessage(), logs.Tags{"error": err})
		return err
	}

	c.JSON(http.StatusOK, response)
	return nil
}
