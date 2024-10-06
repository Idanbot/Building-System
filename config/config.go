package config

type Config struct {
	DB         string
	ServerPort string
}

func Load() *Config {
	return &Config{
		DB:         "postgres://user:password@localhost:5432/building-system?sslmode=disable", //os.Getenv("DB_CONNECTION_STRING"),
		ServerPort: ":5300",
	}
}
