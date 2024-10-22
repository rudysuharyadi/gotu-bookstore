package services

import (
	"bytes"
	"gotu-bookstore/pkg/adapter/api"
	"gotu-bookstore/pkg/adapter/constant"
	"gotu-bookstore/pkg/adapter/templates"
	"gotu-bookstore/pkg/utils"
	"html/template"
)

type ForgotPasswordEmailService struct {
	context    utils.CommonContext
	mailgunAPI api.MailgunAPI
}

func NewForgotPasswordEmailService(context utils.CommonContext, mailgunAPI api.MailgunAPI) ForgotPasswordEmailService {
	return ForgotPasswordEmailService{
		context:    context,
		mailgunAPI: mailgunAPI,
	}
}

func (s ForgotPasswordEmailService) SendEmail(to, token, baseUrl string) error {
	// Define the template for the model code
	tmpl, err := template.
		New(constant.ForgotPasswordTemplateName).
		Parse(templates.ForgotPasswordTemplates)
	if err != nil {
		return err
	}

	// Create a buffer to hold the rendered template
	var resultBuffer bytes.Buffer

	// Execute the template with the provided data and write the result to the buffer
	err = tmpl.Execute(&resultBuffer, templates.ForgotPasswordData{
		Email:   to,
		Token:   token,
		BaseURL: baseUrl,
	})
	if err != nil {
		return err
	}

	response, err := s.mailgunAPI.SendEmail(
		constant.NoReplySender,
		to,
		constant.ForgotPasswordSubject,
		resultBuffer.String(),
	)
	if err != nil {
		return err
	}

	s.context.LogInfo("Email sent with messageId: ", response.Id)
	return nil
}
