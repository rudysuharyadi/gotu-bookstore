package httpclient

import (
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/hystrix"
	"github.com/gojek/heimdall/v7/plugins"
	"net/http"
	"time"
)

type Client struct {
	httpClient *hystrix.Client
}

// Leave it empty to default.
type Option struct {
	Timeout                time.Duration         // Default: 30s
	RetryCount             int                   // Default: 0
	RetryDelayInBetween    time.Duration         // Default: no retrier
	CommandName            string                // Default: empty. This is required to get hystrix function under the same name.
	MaxConcurrentRequest   int                   // Default: 100
	ErrorPercentThreshold  int                   // Default: 25
	SleepWindow            int                   // Default: 10
	RequestVolumeThreshold int                   // Default: 10
	FallbackFn             func(err error) error // Default: empty
	RequestLoggerPlugin    bool                  // Default: false
}

func NewClient(x Option) *Client {
	options := make([]hystrix.Option, 0)

	// Adding fallback function
	fallbackFn := func(err error) error {
		if x.FallbackFn != nil {
			x.FallbackFn(err)
		}
		return err
	}
	options = append(options, hystrix.WithFallbackFunc(fallbackFn))

	// Create retry mechanism.
	if x.RetryCount > 0 {
		backoffInterval := x.RetryDelayInBetween
		maximumJitterInterval := 5 * time.Millisecond
		backoff := heimdall.NewConstantBackoff(backoffInterval, maximumJitterInterval)
		retrier := heimdall.NewRetrier(backoff)
		options = append(options, hystrix.WithRetrier(retrier))
		options = append(options, hystrix.WithRetryCount(x.RetryCount))
	}

	// Adding timeout
	if x.Timeout > 0 {
		options = append(options, hystrix.WithHTTPTimeout(x.Timeout))
		options = append(options, hystrix.WithHystrixTimeout(x.Timeout))
	}

	// Adding CommandName
	if len(x.CommandName) > 0 {
		options = append(options, hystrix.WithCommandName(x.CommandName))
	}

	// Adding MaxConcurrentRequest
	if x.MaxConcurrentRequest > 0 {
		options = append(options, hystrix.WithMaxConcurrentRequests(x.MaxConcurrentRequest))
	}

	// Adding ErrorPercentThreshold
	if x.ErrorPercentThreshold > 0 {
		options = append(options, hystrix.WithErrorPercentThreshold(x.ErrorPercentThreshold))
	}

	// Adding SleepWindow
	if x.SleepWindow > 0 {
		options = append(options, hystrix.WithSleepWindow(x.SleepWindow))
	}

	// Adding RequestVolumeThreshold
	if x.RequestVolumeThreshold > 0 {
		options = append(options, hystrix.WithRequestVolumeThreshold(x.RequestVolumeThreshold))
	}

	hystrixClient := hystrix.NewClient(options...)

	if x.RequestLoggerPlugin {
		requestLogger := plugins.NewRequestLogger(nil, nil)
		hystrixClient.AddPlugin(requestLogger)
	}

	client := Client{
		httpClient: hystrixClient,
	}
	return &client
}

func (client Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := client.httpClient.Do(req)
	return resp, err
}
