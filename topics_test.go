package gokafkaesque

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty"
	"github.com/go-test/deep"
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("health.json"))
	}))
	defer apiStub.Close()
	client := &Client{Rest: resty.New().SetRESTMode().SetHostURL(apiStub.URL)}
	r, _ := client.getHealth()
	if r.StatusCode() != 200 {
		t.Errorf("getHealth() expected %v, got %v", 200, r.StatusCode())
	}
	expectedJSON := Health{
		Response: "Ok",
	}
	var data Health
	client.Rest.JSONUnmarshal(r.Body(), &data)
	if diff := deep.Equal(data, expectedJSON); diff != nil {
		t.Errorf("getHealth() expected %#v, got %#v", expectedJSON, data)
	}
}
