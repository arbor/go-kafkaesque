package gokafkaesque

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestHealth(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/health":
			fmt.Fprint(w, fixture("health.json"))
		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}))
	defer apiStub.Close()
	config := NewConfig().SetURL(apiStub.URL).Build()
	client := NewClient(config)
	r, err := client.GetStatus()
	if err != nil {
		t.Errorf("%v", err.Error())
		t.FailNow()
	}
	if r.GetHealth() != "Ok" {
		t.Errorf("r.GetHealth() expected %v, got %v", "Ok", r.GetHealth())
	}
}

func TestGetTopics(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/topics":
			fmt.Fprint(w, fixture("topics.json"))
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
		expectedName        string
		expectedRetentionMs string
		expectedPartitions  int64
	}{
		{"test_kafka-admin-service", "31536000000", 1},
	}
	for _, tt := range data {
		r, err := client.GetTopic(tt.expectedName)
		if err != nil {
			t.Errorf("%v", err.Error())
			t.FailNow()
		}
		if r.Response.Config.RetentionMs != tt.expectedRetentionMs {
			t.Errorf("RetentionMS expected %v, got %v", tt.expectedRetentionMs, r.Response.Config.RetentionMs)
		}
		if r.Response.Partitions != tt.expectedPartitions {
			t.Errorf("Partitions expected %v, got %v", tt.expectedPartitions, r.Response.Partitions)
		}
	}
}

func TestInvalidGetTopicName(t *testing.T) {
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
		expectedName        string
		expectedRetentionMs string
		expectedPartitions  int64
	}{
		{"", "", 0},
	}
	for _, tt := range data {
		_, err := client.GetTopic(tt.expectedName)
		if err == nil {
			t.Errorf("%v", err.Error())
			t.FailNow()
		}
	}
}
