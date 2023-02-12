package controller

import (
	"net/http"
)

// using pseudo heritage to BaseController
type HealthcheckController struct {
	BaseController
}

func NewHealthcheckController() *HealthcheckController {
	return &HealthcheckController{}
}

func (controller HealthcheckController) CheckHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := struct {
			Status string `json:"status"`
		}{
			Status: "Healthy",
		}
		controller.SendOk(w, r, response, nil)
	}
}
