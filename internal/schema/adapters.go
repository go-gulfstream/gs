package schema

const (
	DefaultStreamStorage storageAdapter = iota
	RedisStreamStorageAdapter
	PostgresStreamStorageAdapter
)

const (
	DefaultName                         = "mem"
	RedisStreamStorageAdapterName       = "redis"
	PostgresStreamStorageAdapterName    = "postgres"
	KafkaStreamPublisherAdapterName     = "kafka"
	ConnectorStreamPublisherAdapterName = "wal"
)

const (
	DefaultStreamPublisher publisherAdapter = iota
	KafkaStreamPublisherAdapter
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
	case DefaultStreamStorage:
		return "Memory"
	case PostgresStreamStorageAdapter:
		return "PostgreSQL"
	case RedisStreamStorageAdapter:
		return "Redis"
	}
	return "Unknown"
}

func (a publisherAdapter) String() string {
	switch a {
	case DefaultStreamPublisher:
		return "Memory"
	case KafkaStreamPublisherAdapter:
		return "Kafka"
	case ConnectorStreamPublisherAdapter:
		return "WAL Connector"
	}
	return "Unknown"
}

var StorageAdapters = map[storageAdapter]string{
	DefaultStreamStorage:         DefaultStreamStorage.String(),
	RedisStreamStorageAdapter:    RedisStreamStorageAdapter.String(),
	PostgresStreamStorageAdapter: PostgresStreamStorageAdapter.String(),
}

var PublisherAdapters = map[publisherAdapter]string{
	DefaultStreamPublisher:          DefaultStreamPublisher.String(),
	KafkaStreamPublisherAdapter:     KafkaStreamPublisherAdapter.String(),
	ConnectorStreamPublisherAdapter: ConnectorStreamPublisherAdapter.String(),
}

func (a publisherAdapter) IsKafka() bool {
	return KafkaStreamPublisherAdapter == a
}

func (a publisherAdapter) IsConnector() bool {
	return ConnectorStreamPublisherAdapter == a
}
