package config

import (
	"os"
)

const (
	defaultHost     = "0.0.0.0"
	defaultPort     = "9999"
	defaultLogLevel = "INFO"

	defaultRedisHost = "localhost"
	defaultRedisPort = "6379"
	defaultRedisDB = 0
	defaultRedisPassword = ""
)

var config *Config

type Config struct {
	DecodeThisApiKey string
	Host             string
	Port             string
	LogLevel         string
	Redis            RedisConfig
}

type RedisConfig struct {
	Host     string
	Port     string
	DB       int
	Password string
}

func (c *Config) read() {
	decodeThisApiKey := os.Getenv("DecodeThisApiKey")

	if decodeThisApiKey == "" {
		panic("DecodeThisApiKet env variable is not set.")
	}

	c.DecodeThisApiKey = decodeThisApiKey
	c.Host = _getEnvOr("Host", defaultHost)
	c.Port = _getEnvOr("Port", defaultPort)
	c.LogLevel = _getEnvOr("LogLevel", defaultLogLevel)

	c.Redis.Host = _getEnvOr("RedisHost", defaultRedisHost)
	c.Redis.Port = _getEnvOr("RedisPort", defaultRedisPort)
	c.Redis.DB = defaultRedisDB
	c.Redis.Password = defaultRedisPassword
}

func _getEnvOr(key, _default string) string {
	result := os.Getenv(key)
	if result == "" {
		return _default
	}

	return result
}

func GetConfig() *Config {
	if config != nil {
		return config
	}

	config = &Config{}
	config.read()
	return config
}
