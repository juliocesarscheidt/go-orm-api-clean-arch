package controller

import (
	"net/http"
)

func Healthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		SendOk(w, r, struct {
			status string
		}{
			status: "Healthy",
		}, nil)
	}
}
