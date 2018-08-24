package gokafkaesque

// Topic have information about a Kafka topic.
type Topic struct {
	Meta `json:"response"`
}

// Meta includes Kafka topic config, partitions, replication
// factor and name.
type Meta struct {
	Config            `json:"config"`
	Partitions        int64  `json:"partitions"`
	ReplicationFactor int64  `json:"replicationFactor"`
	Name              string `json:"name"`
}

// Config contains a Kafka topic retention config in ms.
type Config struct {
	RetentionMs  string `json:"retention.ms"`
	SegmentBytes string `json:"segment.bytes"`
}

// Topics is a list of topic.
type Topics []Topic

// Health returns a response of OK.
type Health struct {
	Response string `json:"response"`
}

// ConfigMap contains additionalprop of a Kafka topic. Used to
// updating existing Kafka topic.
type ConfigMap struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
