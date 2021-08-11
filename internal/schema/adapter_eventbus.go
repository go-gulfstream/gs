package schema

const (
	DefaultStreamPublisher publisherAdapter = iota
	KafkaStreamPublisher
	ConnectorStreamPublisher
)

const (
	KafkaStreamPublisherName     = "kafka"
	ConnectorStreamPublisherName = "wal"
)

var PublisherAdapters = map[publisherAdapter]string{
	DefaultStreamPublisher:   DefaultStreamPublisher.String(),
	KafkaStreamPublisher:     KafkaStreamPublisher.String(),
	ConnectorStreamPublisher: ConnectorStreamPublisher.String(),
}

type publisherAdapter int

func (a publisherAdapter) String() string {
	switch a {
	case DefaultStreamPublisher:
		return "Memory"
	case KafkaStreamPublisher:
		return "Kafka"
	case ConnectorStreamPublisher:
		return "WAL Connector"
	}
	return "Unknown"
}
