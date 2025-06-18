package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `env:"APP_ENV" env-default:"prod"`
	HTTPServer
	Db
}

type HTTPServer struct {
	Host         string        `env:"HTTP_HOST" env-default:"0.0.0.0:8082"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" env-default:"8s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" env-default:"8s"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" env-default:"60s"`
	User         string        `env:"HTTP_USER"`            // env-required:"true"`
	Password     string        `env:"HTTP_SERVER_PASSWORD"` // env-required:"true"`
}

type Db struct {
	DBHost  string `env:"DB_HOST" env-default:"localhost"`
	DBPort  string `env:"DB_PORT"  env-default:"5432"`
	DBUser  string `env:"DB_USER" env-required:"true"`
	DBPass  string `env:"DB_PASSWORD" env-required:"true"`
	DBName  string `env:"DB_NAME" env-required:"true"`
	Sslmode string `env:"SSL_MODE" env-default:"disable"` //disable, require, verify-full

}

func MustLoadEnvConfig() *Config {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Cann't read config enviroment:%v", err)
	}
	return &cfg
}
