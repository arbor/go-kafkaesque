// +build !doc

package gokafkaesque_test

import (
	"fmt"

	gokafkaesque "github.com/packetloop/go-kafkaesque"
)

func ExampleNewClient() {
	config := gokafkaesque.NewConfig().
		SetHost("localhost").
		SetPort(8080).
		SetRetry(3).
		Build()
	client := gokafkaesque.NewClient(config)
	a, _ := client.Rest.R().Get("/health")
	fmt.Println(a)

	// Output:
	// {"response":"Ok"}
}
