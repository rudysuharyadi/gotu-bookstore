package api

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"time"

	"gotu-bookstore/pkg/httpclient"
)

func InitClient(config MailgunConfig) httpclient.Client {
	option := httpclient.Option{
		Timeout:             time.Duration(config.Timeout) * time.Second,
		RetryCount:          config.RetryCount,
		CommandName:         string(constants.HttpclientNameMailgun),
		FallbackFn:          fallback,
		RequestLoggerPlugin: true,
	}
	return *httpclient.NewClient(option)
}

func fallback(err error) error {
	return err
}
