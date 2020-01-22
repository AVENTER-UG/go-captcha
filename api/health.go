package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *Service) apiHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	logrus.Debug("Health Check")

	_, err := s.Cache.Ping().Result()

	if err != nil {
		logrus.Error("Health to Redis Server: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
