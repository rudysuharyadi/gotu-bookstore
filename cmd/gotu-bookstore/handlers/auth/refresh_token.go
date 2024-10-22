package auth

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type RefreshTokenServiceInterface interface {
	ProcessingRefreshToken(request contracts.RefreshTokenRequest) (*contracts.RefreshTokenResponse, error)
}

type RefreshTokenHandler struct {
	context             utils.CommonContext
	refreshTokenService RefreshTokenServiceInterface
}

func NewRefreshTokenHandler(context utils.CommonContext, refreshTokenService RefreshTokenServiceInterface) RefreshTokenHandler {
	return RefreshTokenHandler{
		context:             context,
		refreshTokenService: refreshTokenService,
	}
}

func (h RefreshTokenHandler) ProcessingRefreshToken() {
	var request contracts.RefreshTokenRequest
	if err := h.context.GinContext.BindJSON(&request); err != nil {
		h.context.LogError(err)
		h.context.HandleFailure(base_error.NewBadRequestError(constants.IC0001))
		return
	}

	result, err := h.refreshTokenService.ProcessingRefreshToken(request)
	if err != nil {
		h.context.HandleFailure(err)
		return
	}

	h.context.HandleSuccess(result)
}
