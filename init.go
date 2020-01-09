package main

import (
	"os"

	cfg "./types"
)

var config cfg.Config

func init() {
	config.APIBind = os.Getenv("API_BIND")
	config.APIPort = os.Getenv("API_PORT")
	config.LogLevel = os.Getenv("LOGLEVEL")
	config.RedisServer = os.Getenv("REDIS_SERVER")
}
