package main

import (
	"log"
	"net/http"
	"regexp"
)

type Response struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func jsonContent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func isDigit(value string) bool {
	isNum := regexp.MustCompile(`^\d+$`)

	return isNum.MatchString(value)
}

// / pingRender Keep pinging render to prevent sleeping
func pingRender() {
	_, err := http.Get("https://personapi-l9ah.onrender.com/api")
	if err != nil {
		log.Println(err)
	}
}
