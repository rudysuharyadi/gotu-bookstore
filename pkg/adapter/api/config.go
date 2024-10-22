package api

type MailgunConfig struct {
	BaseURL    string
	DomainName string
	Timeout    int
	RetryCount int
	ApiKey     string
}
