package auth

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type LoginServiceInterface interface {
	ProcessingLogin(request contracts.LoginRequest) (*contracts.LoginResponse, error)
}

type LoginHandler struct {
	context      utils.CommonContext
	loginService LoginServiceInterface
}

func NewLoginHandler(context utils.CommonContext, loginService LoginServiceInterface) LoginHandler {
	return LoginHandler{
		context:      context,
		loginService: loginService,
	}
}

func (h LoginHandler) ProcessingLogin() {
	var request contracts.LoginRequest
	if err := h.context.GinContext.BindJSON(&request); err != nil {
		h.context.LogError(err)
		h.context.HandleFailure(base_error.NewBadRequestError(constants.IC0001))
		return
	}

	result, err := h.loginService.ProcessingLogin(request)
	if err != nil {
		h.context.HandleFailure(err)
		return
	}

	h.context.HandleSuccess(result)
}
