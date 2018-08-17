package gokafkaesque

import (
	"strconv"

	"github.com/go-resty/resty"
)

// Client contains Singularity endpoint for http requests
type Client struct {
	Rest *resty.Client
}

// serverConfig contains Kafka-admin-service HTTP endpoint and serverConfiguration for
// retryablehttp client's retry options
type serverConfig struct {
	Host  string
	Port  int
	Retry int
}

type ServerConfigBuilder interface {
	SetPort(int) ServerConfigBuilder
	SetHost(string) ServerConfigBuilder
	SetRetry(int) ServerConfigBuilder
	Build() serverConfig
}

// NewConfig returns an empty ServerConfigBuilder.
func NewConfig() ServerConfigBuilder {
	return &serverConfig{}
}

// SetHost accepts a string and sets the host in serverConfig.
func (co *serverConfig) SetHost(host string) ServerConfigBuilder {
	co.Host = host
	return co
}

// SetHost accepts an int and sets the retry count.
func (co *serverConfig) SetRetry(r int) ServerConfigBuilder {
	co.Retry = r
	return co
}

// SetHost accepts an int and sets the port number.
func (co *serverConfig) SetPort(port int) ServerConfigBuilder {
	co.Port = port
	return co
}

// Build method returns a serverConfig struct.
func (co *serverConfig) Build() serverConfig {

	return serverConfig{
		Host:  co.Host,
		Port:  co.Port,
		Retry: co.Retry,
	}
}

// NewClient returns Singularity HTTP endpoint.
func NewClient(c serverConfig) *Client {
	r := resty.New().
		SetRESTMode().
		SetRetryCount(c.Retry).
		SetHostURL(endpoint(&c))
	return &Client{
		Rest: r,
	}
}

func endpoint(c *serverConfig) string {
	// if port is uninitialised, port would be http/80.
	if c.Port == 0 || c.Port == 80 {
		return "http://" + c.Host
	}
	if c.Port == 443 {
		return "https://" + c.Host
	}
	return "http://" + c.Host + ":" + strconv.Itoa(c.Port)
}
