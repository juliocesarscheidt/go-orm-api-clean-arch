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
	value, err := strconv.ParseInt(function(field), 10, 64)
	if err != nil {
		return 0, err
	}
	utils.Logger.Infof("Value :: %d", value)
	return int(value), nil
}

func ThrowInternalServerError(w http.ResponseWriter, message string) {
	fmt.Println(message)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func ThrowBadRequest(w http.ResponseWriter, message string) {
	fmt.Println(message)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func ThrowUnauthorized(w http.ResponseWriter, message string) {
	fmt.Println(message)
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func ThrowNotFound(w http.ResponseWriter, message string) {
	fmt.Println(message)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&dto.HttpResponseMessageDto{Message: message})
}

func HandleError(w http.ResponseWriter, err error) {
	fmt.Println(fmt.Sprintf("Message %v", err.Error()))

	if err.Error() == "Not Found" {
		ThrowNotFound(w, err.Error())
		return

	} else if err.Error() == "Bad Request" {
		ThrowBadRequest(w, err.Error())
		return

	} else if err.Error() == "Internal Server Error" {
		ThrowInternalServerError(w, err.Error())
		return
	}

	ThrowInternalServerError(w, err.Error())
}
