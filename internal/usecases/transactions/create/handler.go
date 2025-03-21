package create

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiErrors "github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	apiContext "github.com/melegattip/financial-resume-engine/internal/infrastructure/context"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
)

const createTransaction = "createTransaction"

type CreateTransaction struct {
	Service Create
}

func (handler *CreateTransaction) Handle(c *gin.Context) error {
	ctx := apiContext.SetAction(c.Request.Context(), createTransaction)

	var createRequest CreateTransactionRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		logger.Error(ctx, err, err.Error(), logs.Tags{})
		return apiErrors.NewBadRequest(logs.ErrorBinding.GetMessageWithMapParams(logs.Params{"error": err}))
	}

	userID := c.GetHeader("x-caller-id")
	if userID != createRequest.UserID {
		logger.Error(ctx, nil, logs.ErrorUnauthorizedRequest.GetMessageWithMapParams(
			logs.Params{"x-caller-id": userID}),
			logs.Tags{"user_id": createRequest.UserID})

		return apiErrors.NewUnauthorizedRequest(logs.ErrorUnauthorizedRequest.GetMessage())
	}

	serviceResponse, err := handler.Service.Execute(ctx, &createRequest)
	if err != nil {
		logger.Error(ctx, err, logs.ErrorCreatingTransaction.GetMessage(), logs.Tags{"error": err})
		return err
	}

	c.JSON(http.StatusCreated, serviceResponse)
	return nil
}
