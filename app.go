package main

import (
	"net/http"
	_ "net/http/pprof"

	api "./api"
	util "git.aventer.biz/AVENTER/util"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var MinVersion string

// init the redis cache
func initCache() {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisServer,
		DB:   1,
	})

	pong, err := client.Ping().Result()
	logrus.Debug(pong, err)

	config.Cache = client
}

func main() {
	util.SetLogging(config.LogLevel, config.EnableSyslog, config.AppName)
	logrus.Println(config.AppName+" build"+config.MinVersion, config.APIBind, config.APIPort)

	var s api.Service

	initCache()
	s.SetConfig(config)

	http.Handle("/", s.Commands())

	if err := http.ListenAndServe(config.APIBind+":"+config.APIPort, nil); err != nil {
		logrus.Fatalln("ListenAndServe: ", err)
	}
}
