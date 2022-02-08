package api

import (
	//"encoding/json"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"

	//"io/ioutil"

	"net/http"

	cfg "github.com/AVENTER-UG/go-captcha/types"
	util "github.com/AVENTER-UG/util"
)

type Service struct {
	Cache *redis.Client
}

// SetConfig set the name of the database
func (s *Service) SetConfig(config cfg.Config) {
	s.Cache = config.Cache

	util.SetLogging(config.LogLevel, config.EnableSyslog, config.AppName)
}

// Commands function export all the the api calls
func (s *Service) Commands() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/versions", s.apiVersions).Methods("GET")
	rtr.HandleFunc("/api", s.apiVersions).Methods("GET")
	rtr.HandleFunc("/health", s.apiHealth).Methods("GET")
	rtr.HandleFunc("/api/captcha/v0/version", s.apiV0Version).Methods("GET")
	rtr.HandleFunc("/api/captcha/v0", s.apiV0CaptchaGet).Methods("GET")
	rtr.HandleFunc("/api/captcha/v0/{captcha}", s.apiV0CaptchaCheck).Methods("POST")

	return rtr
}

func (s *Service) apiVersions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "-")
	w.Write([]byte("/api/auth/v0"))
}

func (s *Service) apiV0Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "v0")
	w.Write([]byte("v0.1"))
}
