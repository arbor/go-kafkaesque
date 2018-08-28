package gokafkaesque

import (
	"fmt"
)

// GetTopics is a method that returns all Kafka topics.
func (client *Client) GetTopics() (Topics, error) {
	//func (client *Client) getHealth() string {
	//return client.Rest.HostURL
	resp, err := client.Rest.R().Get("/topics")
	if err != nil {
		return Topics{}, fmt.Errorf("ERROR: %s", err.Error())
	}
	if resp.StatusCode() >= 200 && resp.StatusCode() <= 299 {
		var data Topics
		err := client.Rest.JSONUnmarshal(resp.Body(), &data)
		if err != nil {
			return Topics{}, fmt.Errorf("ERROR: %s", err.Error())
		}
		return data, nil
	}
	return Topics{}, fmt.Errorf("get status error: %v", resp.Status())
}

// Count is a method that returns total size of topics.
func (t Topics) Count() int {
	return len(t.Response.Topics)
}

// Topics is a method that returns a slice of topics.
func (t Topics) Topics() []string {
	return t.Response.Topics
}

// GetTopic is a method that return a Kafka topic
func (client *Client) GetTopic(topic string) (Topic, error) {
	if len(topic) > 0 {
		resp, err := client.Rest.R().Get("/topics/" + topic)
		if err != nil {
			return Topic{}, fmt.Errorf("ERROR: %s", err.Error())
		}
		if resp.StatusCode() >= 200 && resp.StatusCode() <= 299 {
			var data Topic
			err := client.Rest.JSONUnmarshal(resp.Body(), &data)
			if err != nil {
				return Topic{}, fmt.Errorf("ERROR: %s", err.Error())
			}
			return data, nil
		}
	}
	return Topic{}, fmt.Errorf("Please provide a topic name\n")
}
