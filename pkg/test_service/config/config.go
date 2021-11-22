package config

type Config struct {
	Host     string
	GrpcPort int
	RestPort int
}

func New() *Config {

	cfg := Config{
		Host:     "localhost",
		GrpcPort: 8080,
		RestPort: 8081,
	}
	return &cfg
}
