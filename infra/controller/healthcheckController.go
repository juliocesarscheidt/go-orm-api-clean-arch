package controller

import (
	"net/http"
)

type HealthcheckController struct {
}

func NewHealthcheckController() *HealthcheckController {
	return &HealthcheckController{}
}

func (controller HealthcheckController) CheckHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		SendOk(w, r, struct {
			status string
		}{
			status: "Healthy",
		}, nil)
	}
}
