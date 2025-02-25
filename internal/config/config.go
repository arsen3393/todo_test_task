package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig
	DBConfig DataBaseConfig
}

type ServerConfig struct {
	Port           int           `env:"SERVER_PORT" envDefault:"8080"`
	Timeout        time.Duration `env:"SERVER_TIMEOUT" envDefault:"5s"`
	IdleTimeout    time.Duration `env:"SERVER_IDLE_TIMEOUT" envDefault:"60s"`
	SwaggerEnabled bool          `env:"SERVER_SWAGGER_ENABLED" envDefault:"true"`
}

type DataBaseConfig struct {
	Host     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	Name     string `env:"DB_NAME" envDefault:"dev"`
	User     string `env:"DB_USER" envDefault:"root"`
	Password string `env:"DB_PASSWORD" envDefault:"123456"`
	TimeZone string `env:"DB_TIME_ZONE" env-default:"UTC" comment:"Часовой пояс базы данных"`
}

func MustLoadConfig() *Config {
	configPath := ".env"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file not exist")
	}
	cfg := &Config{}
	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		panic("cannot load config file: " + err.Error())
	}
	return cfg
}
