package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/juliocesarscheidt/go-orm-api/application/dto"
	"github.com/juliocesarscheidt/go-orm-api/shared/utils"
)

func GetValueFromParamsAsInt(vars map[string]string, field string) (int, error) {
	value, err := strconv.Atoi(vars[field])
	if err != nil {
		return 0, err
	}
	utils.Logger.Infof("Value :: %d", value)
	return int(value), nil
}

func GetValueFromFormAsInt(function func(key string) string, field string) (int, error) {
	value, err := strconv.Atoi(function(field))
	if err != nil {
		return 0, err
	}
	utils.Logger.Infof("Value :: %d", value)
	return int(value), nil
}

func putMetricsAndSetHeader(w http.ResponseWriter, r *http.Request, statusCode int) {
	utils.PutPrometheusMetrics(r.RequestURI, r.Method, fmt.Sprint(statusCode))
	w.WriteHeader(statusCode)
}

func SendOk(w http.ResponseWriter, r *http.Request, data, metadata interface{}) {
	putMetricsAndSetHeader(w, r, http.StatusOK)
	json.NewEncoder(w).Encode(&dto.HttpResponseDto{Data: data, Metadata: metadata})
}

func SendCreated(w http.ResponseWriter, r *http.Request, data, metadata interface{}) {
	putMetricsAndSetHeader(w, r, http.StatusCreated)
	json.NewEncoder(w).Encode(&dto.HttpResponseDto{Data: data, Metadata: metadata})
}

func SendAccepted(w http.ResponseWriter, r *http.Request, message string) {
	putMetricsAndSetHeader(w, r, http.StatusAccepted)
}

func SendNoContent(w http.ResponseWriter, r *http.Request) {
	putMetricsAndSetHeader(w, r, http.StatusNoContent)
}

func SendBadRequest(w http.ResponseWriter, r *http.Request, message string) {
	putMetricsAndSetHeader(w, r, http.StatusBadRequest)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func SendUnprocessableEntity(w http.ResponseWriter, r *http.Request, message string) {
	putMetricsAndSetHeader(w, r, http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func SendUnauthorized(w http.ResponseWriter, r *http.Request, message string) {
	putMetricsAndSetHeader(w, r, http.StatusUnauthorized)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func SendNotFound(w http.ResponseWriter, r *http.Request, message string) {
	putMetricsAndSetHeader(w, r, http.StatusNotFound)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func SendInternalServerError(w http.ResponseWriter, r *http.Request, message string) {
	putMetricsAndSetHeader(w, r, http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}
