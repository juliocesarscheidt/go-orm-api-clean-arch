package controller

import (
	"net/http"
)

// using composition to BaseController
type HealthcheckController struct {
	BaseController
}

func NewHealthcheckController() *HealthcheckController {
	return &HealthcheckController{}
}

func (controller HealthcheckController) CheckLiveness() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := struct {
			Status string `json:"status"`
		}{
			Status: "Alive",
		}
		controller.SendOk(w, r, response, nil)
	}
}

func (controller HealthcheckController) CheckReadiness() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := struct {
			Status string `json:"status"`
		}{
			Status: "Ready",
		}
		controller.SendOk(w, r, response, nil)
	}
}
