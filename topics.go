package gokafkaesque

import (
	"fmt"

	"github.com/go-resty/resty"
)

func (client *Client) getHealth() (*resty.Response, error) {
	//func (client *Client) getHealth() string {
	//return client.Rest.HostURL
	resp, err := client.Rest.R().Get("/")
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	return resp, nil
}
