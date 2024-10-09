package config

type Config struct {
	DatabaseType   string
	PostgresConfig PostgresConfig
	MongoConfig    MongoConfig
}

type PostgresConfig struct {
	// Параметры подключения к PostgreSQL
}

type MongoConfig struct {
	// Параметры подключения к MongoDB
}
