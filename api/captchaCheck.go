package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// This funktion will check if the captcha is valid
func (s *Service) apiV0CaptchaCheck(w http.ResponseWriter, r *http.Request) {

	jsonStr := []byte(`{"valid": "false"}`)

	sessionToken := r.Header.Get("sessionToken")

	if sessionToken == "" {
		logrus.Error("Session Token is empty")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	if vars == nil {
		logrus.Error("No captcha")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	captcha, err := s.Cache.Get(sessionToken).Result()
	if err != nil {
		logrus.Error("Could not find session token in cache")
		logrus.Debug("Could not find session token in cache: ", sessionToken)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.Debug(sessionToken, vars["captcha"], captcha)

	if captcha != vars["captcha"] {
		logrus.Error("Wrong captcha")
		logrus.Debug("Wrong captcha: ", vars["captcha"], captcha)
		w.Write(jsonStr)
		return
	}

	jsonStr = []byte(`{"valid": "true"}`)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.Write(jsonStr)
}
