package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/shared/utils"
)

type BaseController struct {
}

func (controller BaseController) GetValueFromParamsAsInt(vars map[string]string, field string) (int, error) {
	value, err := strconv.Atoi(vars[field])
	if err != nil {
		return 0, err
	}
	return int(value), nil
}

func (controller BaseController) GetValueFromFormAsInt(function func(key string) string, field string) (int, error) {
	value, err := strconv.Atoi(function(field))
	if err != nil {
		return 0, err
	}
	return int(value), nil
}

func (controller BaseController) SendOk(w http.ResponseWriter, r *http.Request, data, metadata interface{}) {
	utils.PutEndpointMetrics(r.RequestURI, r.Method, fmt.Sprint(http.StatusOK))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&dto.HttpResponseDto{Data: data, Metadata: metadata})
}

func (controller BaseController) SendCreated(w http.ResponseWriter, r *http.Request, data, metadata interface{}) {
	utils.PutEndpointMetrics(r.RequestURI, r.Method, fmt.Sprint(http.StatusCreated))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&dto.HttpResponseDto{Data: data, Metadata: metadata})
}

func (controller BaseController) SendAccepted(w http.ResponseWriter, r *http.Request, message string) {
	utils.PutEndpointMetrics(r.RequestURI, r.Method, fmt.Sprint(http.StatusAccepted))
	w.WriteHeader(http.StatusAccepted)
}

func (controller BaseController) SendNoContent(w http.ResponseWriter, r *http.Request) {
	utils.PutEndpointMetrics(r.RequestURI, r.Method, fmt.Sprint(http.StatusNoContent))
	w.WriteHeader(http.StatusNoContent)
}

func (controller BaseController) SendBadRequest(w http.ResponseWriter, r *http.Request, message string) {
	utils.PutEndpointMetrics(r.RequestURI, r.Method, fmt.Sprint(http.StatusBadRequest))
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func (controller BaseController) SendUnprocessableEntity(w http.ResponseWriter, r *http.Request, message string) {
	utils.PutEndpointMetrics(r.RequestURI, r.Method, fmt.Sprint(http.StatusUnprocessableEntity))
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func (controller BaseController) SendUnauthorized(w http.ResponseWriter, r *http.Request, message string) {
	utils.PutEndpointMetrics(r.RequestURI, r.Method, fmt.Sprint(http.StatusUnauthorized))
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func (controller BaseController) SendNotFound(w http.ResponseWriter, r *http.Request, message string) {
	utils.PutEndpointMetrics(r.RequestURI, r.Method, fmt.Sprint(http.StatusNotFound))
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func (controller BaseController) SendInternalServerError(w http.ResponseWriter, r *http.Request, message string) {
	utils.PutEndpointMetrics(r.RequestURI, r.Method, fmt.Sprint(http.StatusInternalServerError))
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func (controller BaseController) HandleError(w http.ResponseWriter, r *http.Request, err error) {
	message := err.Error()
	utils.Logger.Infof(message)
	if message == "Not found" {
		controller.SendNotFound(w, r, message)
		return
	} else if message == "Bad request" {
		controller.SendBadRequest(w, r, message)
		return
	} else if message == "Invalid password length" {
		controller.SendUnprocessableEntity(w, r, "Invalid password length, the password must have at least 8 and at most 50 characters")
		return
	} else if strings.Contains(message, "Invalid name") ||
		strings.Contains(message, "Invalid email") ||
		strings.Contains(message, "Invalid password") {
		controller.SendBadRequest(w, r, message)
		return
	} else if message == "Internal server error" {
		controller.SendInternalServerError(w, r, message)
		return
	}
	controller.SendInternalServerError(w, r, message)
}
