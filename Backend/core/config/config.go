package config

import "os"

type Config struct {
    DbHost     string
    DbPort     string
    DbUser     string
    DbPassword string
    DbName     string
    JwtSecret  string
    ServerHost string
}

func LoadConfig() (*Config, error){
    return &Config{
        DbHost:     os.Getenv("DB_HOST"),
        DbPort:     os.Getenv("DB_PORT"),
        DbUser:     os.Getenv("DB_USER"),
        DbPassword: os.Getenv("DB_PASSWORD"),
        DbName:     os.Getenv("DB_NAME"),
        JwtSecret:  os.Getenv("JWT_SECRET"),
        ServerHost: os.Getenv("SERVER_HOST"),
    }, nil
}
