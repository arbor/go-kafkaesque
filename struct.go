package gokafkaesque

// TopicResponse have information about a Kafka topic.
type TopicResponse struct {
	Topic `json:"response"`
}

// Topic includes Kafka topic config, partitions, replication
// factor and name.
type Topic struct {
	*Config           `json:"config"`
	Partitions        int64   `json:"partitions"`
	ReplicationFactor int64   `json:"replicationFactor"`
	Name              *string `json:"name"`
}

// Config contains a Kafka topic retention config in ms.
type Config struct {
	RetentionMs  string `json:"retention.ms"`
	SegmentBytes string `json:"segment.bytes"`
}

// Topics is a list of topic names.
type Topics struct {
	Response struct {
		Topics []string `json:"topics"`
	} `json:"response"`
}

// GenericResponse returns a response of OK.
type GenericResponse struct {
	Response string `json:"response"`
}

// ConfigMap contains additionalprop of a Kafka topic. Used to
// updating existing Kafka topic.
type ConfigMap struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
