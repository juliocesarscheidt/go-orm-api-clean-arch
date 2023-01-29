package controller

import (
	"encoding/json"
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
	value, err := strconv.ParseInt(function(field), 10, 64)
	if err != nil {
		return 0, err
	}
	utils.Logger.Infof("Value :: %d", value)
	return int(value), nil
}

func ThrowInternalServerError(w http.ResponseWriter, message string) {
	utils.Logger.Infof(message)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func ThrowBadRequest(w http.ResponseWriter, message string) {
	utils.Logger.Infof(message)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func ThrowUnauthorized(w http.ResponseWriter, message string) {
	utils.Logger.Infof(message)
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func ThrowNotFound(w http.ResponseWriter, message string) {
	utils.Logger.Infof(message)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func HandleError(w http.ResponseWriter, err error) {
	message := err.Error()
	utils.Logger.Infof(message)

	if message == "Not Found" {
		ThrowNotFound(w, message)
		return
	} else if message == "Bad Request" {
		ThrowBadRequest(w, message)
		return
	} else if message == "Invalid Name" || message == "Invalid Email" || message == "Invalid Password" {
		ThrowBadRequest(w, message)
		return
	} else if message == "Internal Server Error" {
		ThrowInternalServerError(w, message)
		return
	}

	ThrowInternalServerError(w, message)
}
