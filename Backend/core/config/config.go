package config

import "os"

type Config struct {
    DatabaseURL string
    ServerAddress string
}

func LoadConfig() (*Config, error) {
    return &Config{
        DatabaseURL:   os.Getenv("DATABASE_URL"),
        ServerAddress: os.Getenv("SERVER_ADDRESS"),
    }, nil
}