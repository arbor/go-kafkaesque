package gokafkaesque

import (
	"testing"

	"gopkg.in/resty.v1"
)

func TestBuild(t *testing.T) {

	var data = []struct {
		url            string
		retry          int
		expectedConfig serverConfig
	}{
		{"http://localhost", 2, serverConfig{URL: "http://localhost", Retry: 2}},
		{"https://localhost2", 1, serverConfig{URL: "https://localhost2", Retry: 1}},
		{"http://localhost:8080", 0, serverConfig{URL: "http://localhost:8080", Retry: 0}},
	}

	for _, tt := range data {
		c := NewConfig().SetURL(tt.url).SetRetry(tt.retry).Build()
		if c != tt.expectedConfig {
			t.Errorf("NewConfig().SetURL(%s).SetRetry(%d).Build(): expected %v, got %v",
				tt.url,
				tt.retry,
				tt.expectedConfig,
				c)
		}
	}
}

func TestSetURL(t *testing.T) {
	var data = []struct {
		url         string
		expectedURL string
	}{
		{"http://localhost", "http://localhost"},
		{"myendpoint", "myendpoint"},
	}

	for _, tt := range data {
		c := NewConfig().SetURL(tt.url).Build()
		if c.URL != tt.expectedURL {
			t.Errorf("NewConfig().SetHost(%s).Build(): expected %v, got %v",
				tt.url,
				tt.expectedURL,
				c.URL)
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
		url            string
		retry          int
		expectedClient *Client
	}{
		{"http://localhost", 2, &Client{Rest: &resty.Client{
			RetryCount: 2,
			HostURL:    "http://localhost",
		}}},
		{"http://localhost2:8080", 4, &Client{Rest: &resty.Client{
			RetryCount: 4,
			HostURL:    "http://localhost2:8080",
		}}},
	}

	for _, tt := range data {
		config := NewConfig().SetURL(tt.url).SetRetry(tt.retry).Build()
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
