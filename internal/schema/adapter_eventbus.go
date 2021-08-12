package schema

const (
	KafkaStreamPublisher publisherAdapter = iota
	ConnectorStreamPublisher
)

const (
	KafkaStreamPublisherName     = "kafka"
	ConnectorStreamPublisherName = "wal"
)

var PublisherAdapters = map[publisherAdapter]string{
	KafkaStreamPublisher:     KafkaStreamPublisher.String(),
	ConnectorStreamPublisher: ConnectorStreamPublisher.String(),
}

type publisherAdapter int

func (a publisherAdapter) IsKafka() bool {
	return KafkaStreamPublisher == a
}

func (a publisherAdapter) IsConnector() bool {
	return ConnectorStreamPublisher == a
}

func (a publisherAdapter) String() string {
	switch a {
	case KafkaStreamPublisher:
		return "Kafka"
	case ConnectorStreamPublisher:
		return "WAL Connector"
	}
	return "Unknown"
}
