package api

import (
	"bytes"
	"encoding/json"
	"gotu-bookstore/pkg/httpclient"
	"gotu-bookstore/pkg/resfmt/base_error"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

type MailgunAPI struct {
	client httpclient.Client
	config MailgunConfig
}

func NewMailgunAPI(client httpclient.Client, config MailgunConfig) MailgunAPI {
	return MailgunAPI{
		client: client,
		config: config,
	}
}

func (s MailgunAPI) SendEmail(from, to, subject, message string) (*SendEmailResponse, error) {
	reqUrl := s.config.BaseURL + "/" + s.config.DomainName + "/messages"

	data := &bytes.Buffer{}
	writer := multipart.NewWriter(data)
	fromFw, _ := writer.CreateFormField("from")
	_, err := io.Copy(fromFw, strings.NewReader(from+"@"+s.config.DomainName))
	if err != nil {
		return nil, err
	}
	toFw, _ := writer.CreateFormField("to")
	_, err = io.Copy(toFw, strings.NewReader(to))
	if err != nil {
		return nil, err
	}
	subjectFw, _ := writer.CreateFormField("subject")
	_, err = io.Copy(subjectFw, strings.NewReader(subject))
	if err != nil {
		return nil, err
	}
	htmlFw, _ := writer.CreateFormField("html")
	_, err = io.Copy(htmlFw, strings.NewReader(message))
	if err != nil {
		return nil, err
	}
	writer.Close()

	payload := bytes.NewReader(data.Bytes())

	req, err := http.NewRequest("POST", reqUrl, payload)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("api", s.config.ApiKey)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	// If response not ok
	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
		return nil, base_error.New("Error when send email to third party api.")
	}

	// If response OK, read body as []byte, then parse into struct.
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response *SendEmailResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
