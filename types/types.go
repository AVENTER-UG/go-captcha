package types

import (
	"github.com/go-redis/redis"
)

type Config struct {
	APIPort      string
	APIBind      string
	LogLevel     string
	MinVersion   string
	AppName      string
	EnableSyslog bool
	RedisServer  string
	Cache        *redis.Client
}
