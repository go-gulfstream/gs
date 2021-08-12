package schema

const (
	DefaultStreamStorage storageAdapter = iota
	RedisStreamStorage
	PostgresStreamStorage
)

const (
	DefaultName               = "mem"
	RedisStreamStorageName    = "redis"
	PostgresStreamStorageName = "postgres"
)

var StorageAdapters = map[storageAdapter]string{
	DefaultStreamStorage:  DefaultStreamStorage.String(),
	RedisStreamStorage:    RedisStreamStorage.String(),
	PostgresStreamStorage: PostgresStreamStorage.String(),
}

type storageAdapter int

func (a storageAdapter) IsRedis() bool {
	return RedisStreamStorage == a
}

func (a storageAdapter) IsPostgreSQL() bool {
	return PostgresStreamStorage == a
}

func (a storageAdapter) String() string {
	switch a {
	case DefaultStreamStorage:
		return "Memory"
	case PostgresStreamStorage:
		return "PostgreSQL"
	case RedisStreamStorage:
		return "Redis"
	}
	return "Unknown"
}
