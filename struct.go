package gokafkaesque

type Topic struct {
	Meta `json:"response"`
}
type Meta struct {
	Config            `json:"config"`
	Partitions        int64  `json:"partitions"`
	ReplicationFactor int64  `json:"replicationFactor"`
	Name              string `json:"name"`
}

type Config struct {
	Config struct {
		string `json:"retention.ms"`
	}
}

type Topics []Topic

type Health struct {
	Response string `json:"response"`
}

type ConfigMap struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
