package auth

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type RegisterServiceInterface interface {
	ProcessingRegister(request contracts.RegisterRequest) (*contracts.RegisterResponse, error)
}

type RegisterHandler struct {
	registerService RegisterServiceInterface
	utils.CommonContext
}

func NewRegisterHandler(context utils.CommonContext, registerService RegisterServiceInterface) RegisterHandler {
	return RegisterHandler{
		registerService: registerService,
		CommonContext:   context,
	}
}

func (h RegisterHandler) ProcessingRegister() {
	var request contracts.RegisterRequest
	if err := h.GinContext.BindJSON(&request); err != nil {
		h.LogError(err)
		h.HandleFailure(base_error.NewBadRequestError(constants.IC0001))
		return
	}

	result, err := h.registerService.ProcessingRegister(request)
	if err != nil {
		h.HandleFailure(err)
		return
	}

	h.HandleSuccess(result)
}
