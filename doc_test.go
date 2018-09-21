// +build !doc

package gokafkaesque_test

import (
	"fmt"

	gokafkaesque "github.com/packetloop/go-kafkaesque"
)

func ExampleNewClient() {
	config := gokafkaesque.NewConfig().
		SetURL("http://localhost:8080").
		SetRetry(3).
		Build()
	client := gokafkaesque.NewClient(config)
	a, _ := client.GetStatus()
	fmt.Println(a.GetMessage())

	// Output:
	// Ok
}

func ExampleClient_GetTopics() {
	config := gokafkaesque.NewConfig().
		SetURL("http://localhost:8080").
		SetRetry(3).
		Build()
	client := gokafkaesque.NewClient(config)
	a, _ := client.GetTopics()
	fmt.Println(a.TopicsToString())

	// Output:
	// [__confluent.support.metrics __consumer_offsets]
}

func ExampleClient_CreateTopic() {
	config := gokafkaesque.NewConfig().
		SetURL("http://localhost:8080").
		SetRetry(3).
		Build()
	client := gokafkaesque.NewClient(config)
	t := gokafkaesque.NewTopic("foo").SetPartitions("2").SetReplicationFactor("5").BuildTopic()
	a, _ := client.CreateTopic(t)
	fmt.Println(a.GetMessage())

	// output:
	// foo created.
}

func ExampleClient_GetTopic() {
	config := gokafkaesque.NewConfig().
		SetURL("http://localhost:8080").
		SetRetry(3).
		Build()
	client := gokafkaesque.NewClient(config)
	a, _ := client.GetTopic("foo")
	fmt.Printf("partition: %s\nreplication_factor: %s\n", a.GetPartitions(), a.GetReplicationFactor())

	// output:
	// partition: 2
	// replication_factor: 5
}

// This is a PUT request, which means it requires complete parameters set.
// To only update optional params, we need to implement PATCH request instead.
func ExampleClient_UpdateTopic() {
	config := gokafkaesque.NewConfig().
		SetURL("http://localhost:8080").
		SetRetry(3).
		Build()
	client := gokafkaesque.NewClient(config)
	t := gokafkaesque.NewTopic("foo").SetPartitions("2").SetReplicationFactor("5").BuildTopic()
	_, err := client.CreateTopic(t)
	if err != nil {
		fmt.Println(err)
	}
	t.Config = &gokafkaesque.Config{
		RetentionMs:       "1000",
		SegmentBytes:      "10000000",
		CleanupPolicy:     "delete",
		MinInsyncReplicas: "1",
		RetentionBytes:    "10",
		SegmentMs:         "10",
	}
	a, err := client.UpdateTopic(t)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a.GetMessage())

	// output:
	// foo updated.
}
func ExampleClient_DeleteTopic() {
	config := gokafkaesque.NewConfig().
		SetURL("http://localhost:8080").
		SetRetry(3).
		Build()
	client := gokafkaesque.NewClient(config)
	a, _ := client.DeleteTopic("foo")
	fmt.Println(a.GetMessage())

	// Output:
	// foo deleted.
}
