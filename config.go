package gokafkaesque

import (
	"net/http"
	"time"

	"gopkg.in/resty.v1"
)

// Client contains Singularity endpoint for http requests
type Client struct {
	Rest *resty.Client
}

// serverConfig contains Kafka-admin-service HTTP endpoint and serverConfiguration for
// retryablehttp client's retry options
type serverConfig struct {
	URL           string
	Retry         int
	RetryWaitTime time.Duration // Default is 100 milliseconds.
}

// ServerConfigBuilder sets port, host, http retry count config to
// be passed to create a NewClient.
type ServerConfigBuilder interface {
	SetURL(string) ServerConfigBuilder
	SetRetry(int) ServerConfigBuilder
	SetRetryWaitTime(int) ServerConfigBuilder
	Build() serverConfig
}

// NewConfig returns an empty ServerConfigBuilder.
func NewConfig() ServerConfigBuilder {
	return &serverConfig{}
}

// SetHost accepts a string in the form of http://url and sets
// this as URL in serverConfig.
func (co *serverConfig) SetURL(URL string) ServerConfigBuilder {
	co.URL = URL
	return co
}

// SetHost accepts an int and sets the retry count.
func (co *serverConfig) SetRetry(r int) ServerConfigBuilder {
	co.Retry = r
	return co
}

// SetRetryWaitTime accepts an int and sets retry wait time
// in seconds.
func (co *serverConfig) SetRetryWaitTime(t int) ServerConfigBuilder {
	co.RetryWaitTime = time.Duration(t) * time.Second
	return co
}

// Build method returns a serverConfig struct.
func (co *serverConfig) Build() serverConfig {

	return serverConfig{
		URL:           co.URL,
		Retry:         co.Retry,
		RetryWaitTime: co.RetryWaitTime,
	}
}

// NewClient returns Singularity HTTP endpoint.
func NewClient(c serverConfig) *Client {
	r := resty.New().
		SetRESTMode().
		SetRetryCount(c.Retry).
		SetRetryWaitTime(c.RetryWaitTime).
		SetHostURL(c.URL).
		AddRetryCondition(
			// Condition function will be provided with *resty.Response as a
			// parameter. It is expected to return (bool, error) pair. Resty will retry
			// in case condition returns true or non nil error.
			func(r *resty.Response) (bool, error) {
				return r.StatusCode() == http.StatusNotFound, nil
			},
		)
	return &Client{
		Rest: r,
	}
}
