package auth

import (
	"gotu-bookstore/pkg/resfmt/response_format"
	"gotu-bookstore/pkg/utils"
)

type LogoutServiceInterface interface {
	ProcessingLogout() error
}

type LogoutHandler struct {
	context       utils.CommonContext
	logoutService LogoutServiceInterface
}

func NewLogoutHandler(context utils.CommonContext, logoutService LogoutServiceInterface) LogoutHandler {
	return LogoutHandler{
		context:       context,
		logoutService: logoutService,
	}
}

func (h LogoutHandler) ProcessingLogout() {
	err := h.logoutService.ProcessingLogout()
	if err != nil {
		h.context.HandleFailure(err)
		return
	}

	response := response_format.NewSuccess(nil)
	h.context.GinContext.JSON(response.StatusCode, response)
}
