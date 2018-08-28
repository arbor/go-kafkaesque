package gokafkaesque

import (
	"fmt"
)

// GetStatus returns status kafka-admin-service /health endpoint.
func (client *Client) GetStatus() (Health, error) {
	//func (client *Client) getHealth() string {
	//return client.Rest.HostURL
	resp, err := client.Rest.R().Get("/health")
	if err != nil {
		return Health{}, fmt.Errorf("ERROR: %s", err.Error())
	}
	if resp.StatusCode() >= 200 && resp.StatusCode() <= 299 {
		var data Health
		err := client.Rest.JSONUnmarshal(resp.Body(), &data)
		if err != nil {
			return Health{}, fmt.Errorf("ERROR: %s", err.Error())
		}
		return data, nil
	}
	return Health{}, fmt.Errorf("get status error: %v", resp.Status())
}

// GetHealth is a method that returns actual health status of "ok".
func (h *Health) GetHealth() string {
	return h.Response
}
