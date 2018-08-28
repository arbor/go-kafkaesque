package gokafkaesque

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTopics(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/topics":
			fmt.Fprint(w, fixture("listTopics.json"))
		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}))
	defer apiStub.Close()
	config := NewConfig().SetURL(apiStub.URL).Build()
	client := NewClient(config)
	var data = []struct {
		expectedCount int
		expectedName  string
	}{
		{9, "__confluent.support.metrics"},
		{9, "test_kafka-admin-service"},
		{9, "__consumer_offsets"},
	}
	for _, tt := range data {
		r, err := client.GetTopics()
		if err != nil {
			t.Errorf("%v", err.Error())
			t.FailNow()
		}
		if r.Count() != tt.expectedCount {
			t.Errorf("r.Count() expected %v, got %v", tt.expectedCount, r.Count())
		}
		if !contains(r, tt.expectedName) {
			t.Errorf("r.GetTopics() expected %v in %v, got %v", tt.expectedName, r, r)
		}
	}
}

func contains(ts Topics, t string) bool {
	for _, i := range ts.Topics() {
		if i == t {
			return true
		}
	}
	return false
}

func TestGetTopic(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/topics/test_kafka-admin-service":
			fmt.Fprint(w, fixture("topic.json"))
		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}))
	defer apiStub.Close()
	config := NewConfig().SetURL(apiStub.URL).Build()
	client := NewClient(config)
	var data = []struct {
		expectedName              string
		expectedRetentionMs       string
		expectedPartitions        int64
		expectedReplicationFactor int64
	}{
		{"test_kafka-admin-service", "31536000000", 1, 3},
	}
	for _, tt := range data {
		r, err := client.GetTopic(tt.expectedName)
		if err != nil {
			t.Errorf("%v", err.Error())
			t.FailNow()
		}
		if r.Config.GetRetentionMs() != tt.expectedRetentionMs {
			t.Errorf("client.GetTopic(%s) RetentionMS expected %v, got %v", tt.expectedName, tt.expectedRetentionMs, r.Config.GetRetentionMs())
		}
		if r.GetPartitions() != tt.expectedPartitions {
			t.Errorf("client.GetTopic(%s) Partitions expected %v, got %v", tt.expectedName, tt.expectedPartitions, r.GetPartitions())
		}
		if r.GetReplicationFactor() != tt.expectedReplicationFactor {
			t.Errorf("client.GetTopic(%s) ReplicationFactor expected %v, got %v", tt.expectedName, tt.expectedReplicationFactor, r.GetReplicationFactor())
		}
	}
}

func TestInvalidGetTopicName(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/topics/":
			fmt.Fprint(w, fixture("topic.json"))
		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
	}))
	defer apiStub.Close()
	config := NewConfig().SetURL(apiStub.URL).Build()
	client := NewClient(config)
	var data = []struct {
		expectedName     string
		expectedErrorMsg string
	}{
		{"abc", "Status"},
	}
	for _, tt := range data {
		_, err := client.GetTopic(tt.expectedName)
		if err == nil {
			t.Errorf("%v", err.Error())
			t.FailNow()
		}
	}
}

func TestCreateTopic(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/topics":
			fmt.Fprint(w, fixture("createTopic.json"))
		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}))
	defer apiStub.Close()
	config := NewConfig().SetURL(apiStub.URL).Build()
	client := NewClient(config)
	var data = []struct {
		name              string
		partitions        int64
		replicationFactor int64
		expectedResponse  string
	}{
		{"foo", 1, 3, "Ok"},
		{"bar", 1, 3, "Ok"},
	}
	for _, tt := range data {
		params := NewTopic(tt.name).SetPartition(tt.partitions).SetReplicationFactor(tt.replicationFactor).BuildTopic()
		r, err := client.CreateTopic(params)
		if err != nil {
			t.Errorf("%v", err.Error())
			t.FailNow()
		}
		if r.Response != tt.expectedResponse {
			t.Errorf("r.Count() expected %v, got %v", tt.expectedResponse, r.Response)
		}
	}
}

func TestDeleteTopic(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/topics/foo":
			fmt.Fprint(w, fixture("deleteTopic.json"))
		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}))
	defer apiStub.Close()
	config := NewConfig().SetURL(apiStub.URL).Build()
	client := NewClient(config)
	var data = []struct {
		name             string
		expectedResponse string
	}{
		{"foo", "Topic deleted: foo"},
	}
	for _, tt := range data {
		r, err := client.DeleteTopic(tt.name)
		if err != nil {
			t.Errorf("%v", err.Error())
			t.FailNow()
		}
		if r.Response != tt.expectedResponse {
			t.Errorf("client.DeleteTopic(%s) expected %v, got %v", tt.name, tt.expectedResponse, r.Response)
		}
	}
}
