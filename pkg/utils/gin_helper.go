package utils

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/pkg/auth/dto"

	"gotu-bookstore/pkg/logger"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/resfmt/response_format"

	"github.com/gin-gonic/gin"
)

type CommonContext struct {
	GinContext  *gin.Context
	LogInstance *logger.Log
	LogContext  *logger.LogContext
}

func NewCommonContext(ginContext *gin.Context, logInstance *logger.Log, logContext *logger.LogContext) CommonContext {
	return CommonContext{
		GinContext:  ginContext,
		LogInstance: logInstance,
		LogContext:  logContext,
	}
}

// LogInstance with LogContext - Log error with additional context
func (c CommonContext) LogErrors(errs []error, msgs ...string) {
	if c.LogInstance == nil {
		return
	}
	args := make([]interface{}, 0)
	for _, err := range errs {
		args = append(args, err)
	}
	for _, msg := range msgs {
		args = append(args, msg)
	}
	if c.LogContext == nil {
		c.LogInstance.Error(args...)
	} else {
		c.LogInstance.ErrorWithContext(*c.LogContext, args...)
	}
}

func (c CommonContext) LogError(err error, msgs ...string) {
	c.LogErrors([]error{err}, msgs...)
}

func (c CommonContext) LogDebugs(errs []error, msgs ...string) {
	if c.LogInstance == nil {
		return
	}
	args := make([]interface{}, 0)
	for _, err := range errs {
		args = append(args, err)
	}
	for _, msg := range msgs {
		args = append(args, msg)
	}
	if c.LogContext == nil {
		c.LogInstance.Debug(args...)
	} else {
		c.LogInstance.DebugWithContext(*c.LogContext, args...)
	}
}

func (c CommonContext) LogDebug(err error, msgs ...string) {
	c.LogDebugs([]error{err}, msgs...)
}

func (c CommonContext) LogInfo(msgs ...string) {
	if c.LogInstance == nil {
		return
	}
	var args = make([]interface{}, 0)
	for _, msg := range msgs {
		args = append(args, msg)
	}
	if c.LogContext == nil {
		c.LogInstance.Info(args...)
	} else {
		c.LogInstance.InfoWithContext(*c.LogContext, args...)
	}
}

// GinContext - Get objects.
func (c CommonContext) GetSession() (*dto.SessionDTO, error) {
	data, err := c.GetInterfaceFromGinContext(constants.SessionDataContext)
	if err != nil {
		return nil, err
	}
	session, ok := data.(dto.SessionDTO)
	if !ok {
		return nil, base_error.New("Unable to parse object into session")
	}
	return &session, nil
}

func (c CommonContext) GetAccessToken() (string, error) {
	data, err := c.GetInterfaceFromGinContext(constants.AccessTokenContext)
	if err != nil {
		return "", err
	}
	dataString, ok := data.(string)
	if !ok {
		return "", base_error.New("Unable to parse object into string")
	}
	return dataString, nil
}

func (c CommonContext) GetRequestId() (string, error) {
	data, err := c.GetInterfaceFromGinContext(constants.RequestIDContext)
	if err != nil {
		return "", err
	}
	dataString, ok := data.(string)
	if !ok {
		return "", base_error.New("Unable to parse object into string")
	}
	return dataString, nil
}

func (c CommonContext) GetInterfaceFromGinContext(key string) (interface{}, error) {
	if c.GinContext == nil {
		return nil, base_error.New("Gin context not found")
	}
	out, ok := c.GinContext.Get(key)
	if !ok {
		return nil, base_error.New("Key not found")
	}
	return out, nil
}

// GinContext - Handle Failure and Success (with metadata)
func (c CommonContext) HandleFailure(err error) {
	response := response_format.NewFailure(err)
	c.GinContext.Error(err)
	c.GinContext.JSON(response.StatusCode, response)
}

func (c CommonContext) HandleSuccess(result interface{}) {
	response := response_format.NewSuccess(result)
	c.GinContext.JSON(response.StatusCode, response)
}

func (c CommonContext) HandleSuccessWithMetadata(result interface{}, metadata map[string]interface{}) {
	response := response_format.NewSuccessWithMetadata(result, metadata)
	c.GinContext.JSON(response.StatusCode, response)
}
