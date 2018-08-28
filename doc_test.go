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
	fmt.Println(a.GetResponse())

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
	fmt.Println(a.Topics())

	// Output:
	// [__confluent.support.metrics __consumer_offsets]
}

func ExampleClient_CreateTopic() {
	config := gokafkaesque.NewConfig().
		SetURL("http://localhost:8080").
		SetRetry(3).
		Build()
	client := gokafkaesque.NewClient(config)
	t := gokafkaesque.NewTopic("foo").SetPartition(2).SetReplicationFactor(5).BuildTopic()
	a, _ := client.CreateTopic(t)
	fmt.Println(a.GetResponse())

	// output:
	// Ok
}

func ExampleClient_GetTopic() {
	config := gokafkaesque.NewConfig().
		SetURL("http://localhost:8080").
		SetRetry(3).
		Build()
	client := gokafkaesque.NewClient(config)
	a, _ := client.GetTopic("foo")
	fmt.Printf("partition: %d\nreplication_factor: %d\n", a.GetPartitions(), a.GetReplicationFactor())

	// output:
	// partition: 2
	// replication_factor: 5
}

func ExampleClient_DeleteTopic() {
	config := gokafkaesque.NewConfig().
		SetURL("http://localhost:8080").
		SetRetry(3).
		Build()
	client := gokafkaesque.NewClient(config)
	a, _ := client.DeleteTopic("foo")
	fmt.Println(a.GetResponse())

	// Output:
	// Topic deleted: foo
}
