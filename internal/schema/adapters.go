package schema

const (
	RedisStreamStorageAdapter storageAdapter = iota
	PostgresStreamStorageAdapter
)

const (
	KafkaStreamPublisherAdapter publisherAdapter = iota
	ConnectorStreamPublisherAdapter
)

type (
	storageAdapter   int
	publisherAdapter int
)

func (a storageAdapter) IsRedis() bool {
	return RedisStreamStorageAdapter == a
}

func (a storageAdapter) IsPostgreSQL() bool {
	return PostgresStreamStorageAdapter == a
}

func (a storageAdapter) String() string {
	switch a {
	case PostgresStreamStorageAdapter:
		return "PostgreSQL"
	case RedisStreamStorageAdapter:
		return "Redis"
	}
	return "Unknown"
}

func (a publisherAdapter) String() string {
	switch a {
	case KafkaStreamPublisherAdapter:
		return "Kafka"
	case ConnectorStreamPublisherAdapter:
		return "WAL Connector"
	}
	return "Unknown"
}

var StorageAdapters = map[storageAdapter]string{
	RedisStreamStorageAdapter:    RedisStreamStorageAdapter.String(),
	PostgresStreamStorageAdapter: PostgresStreamStorageAdapter.String(),
}

var PublisherAdapters = map[publisherAdapter]string{
	KafkaStreamPublisherAdapter:     KafkaStreamPublisherAdapter.String(),
	ConnectorStreamPublisherAdapter: ConnectorStreamPublisherAdapter.String(),
}

func (a publisherAdapter) IsKafka() bool {
	return KafkaStreamPublisherAdapter == a
}

func (a publisherAdapter) IsConnector() bool {
	return ConnectorStreamPublisherAdapter == a
}
