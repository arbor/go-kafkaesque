package gokafkaesque

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// This is a helper function that returns a JSON response for a request.
func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestHealthMessage(t *testing.T) {
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
	if r.GetMessage() != "Ok" {
		t.Errorf("r.GetMessage() expected %v, got %v", "Ok", r.GetMessage())
	}
}
