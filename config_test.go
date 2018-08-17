package gokafkaesque

import (
	"testing"

	"github.com/go-resty/resty"
)

func TestBuild(t *testing.T) {

	var data = []struct {
		host           string
		port           int
		retry          int
		expectedConfig serverConfig
	}{
		{"localhost", 80, 2, serverConfig{Host: "localhost", Port: 80, Retry: 2}},
		{"localhost2", 443, 1, serverConfig{Host: "localhost2", Port: 443, Retry: 1}},
		{"localhost", 80, 0, serverConfig{Host: "localhost", Port: 80, Retry: 0}},
	}

	for _, tt := range data {
		c := NewConfig().SetHost(tt.host).SetPort(tt.port).SetRetry(tt.retry).Build()
		if c != tt.expectedConfig {
			t.Errorf("NewConfig().SetHost(%s).SetPort(%d).SetRetry(%d).Build(): expected %v, got %v",
				tt.host,
				tt.port,
				tt.retry,
				tt.expectedConfig,
				c)
		}
	}
}

func TestSetPort(t *testing.T) {
	var data = []struct {
		port         int
		expectedPort int
	}{
		{80, 80},
		{8080, 8080},
	}

	for _, tt := range data {
		c := NewConfig().SetPort(tt.port).Build()
		if c.Port != tt.expectedPort {
			t.Errorf("NewConfig().SetPort(%d).Build(): expected %v, got %v",
				tt.port,
				tt.expectedPort,
				c.Port)
		}
	}
}

func TestSetHost(t *testing.T) {
	var data = []struct {
		host         string
		expectedHost string
	}{
		{"localhost", "localhost"},
		{"myendpoint", "myendpoint"},
	}

	for _, tt := range data {
		c := NewConfig().SetHost(tt.host).Build()
		if c.Host != tt.expectedHost {
			t.Errorf("NewConfig().SetHost(%s).Build(): expected %v, got %v",
				tt.host,
				tt.expectedHost,
				c.Host)
		}
	}
}

func TestSetRetry(t *testing.T) {
	var data = []struct {
		retry         int
		expectedRetry int
	}{
		{2, 2},
		{0, 0},
	}

	for _, tt := range data {
		c := NewConfig().SetRetry(tt.retry).Build()
		if c.Retry != tt.expectedRetry {
			t.Errorf("NewConfig().SetRetry(%d).Build(): expected %v, got %v",
				tt.retry,
				tt.expectedRetry,
				c.Retry)
		}
	}

}

func TestNewClient(t *testing.T) {

	var data = []struct {
		host           string
		port           int
		retry          int
		expectedClient *Client
	}{
		{"localhost", 80, 2, &Client{Rest: &resty.Client{
			RetryCount: 2,
			HostURL:    "http://localhost",
		}}},
		{"localhost2", 443, 4, &Client{Rest: &resty.Client{
			RetryCount: 4,
			HostURL:    "https://localhost2",
		}}},
	}

	for _, tt := range data {
		config := NewConfig().SetHost(tt.host).SetPort(tt.port).SetRetry(tt.retry).Build()
		client := NewClient(config)

		if client.Rest.RetryCount != tt.expectedClient.Rest.RetryCount {
			t.Errorf("NewClient(%v): expected %v, got %v",
				&config,
				tt.expectedClient.Rest.RetryCount,
				client.Rest.RetryCount)
		}
		if client.Rest.HostURL != tt.expectedClient.Rest.HostURL {
			t.Errorf("NewClient(%v): expected %v, got %v",
				&config,
				tt.expectedClient.Rest.HostURL,
				client.Rest.HostURL)
		}
	}
}

func TestEndpoint(t *testing.T) {
	c := serverConfig{
		Host: "localhost",
	}
	host := endpoint(&c)
	expected := "http://localhost"
	if host != expected {
		t.Errorf("Got %s, expected %s", host, expected)
	}

	cHTTP := serverConfig{
		Host: "localhost",
		Port: 80,
	}
	hostHTTP := endpoint(&cHTTP)
	expectedHTTP := "http://localhost"
	if hostHTTP != expectedHTTP {
		t.Errorf("Got %s, expected %s", hostHTTP, expectedHTTP)
	}

	cHTTPS := serverConfig{
		Host: "localhost",
		Port: 443,
	}
	hostHTTPS := endpoint(&cHTTPS)
	expectedHTTPS := "https://localhost"
	if hostHTTPS != expectedHTTPS {
		t.Errorf("Got %s, expected %s", hostHTTPS, expectedHTTPS)
	}

}
