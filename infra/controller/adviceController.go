package controller

import (
	"net/http"

	"github.com/juliocesarscheidt/go-orm-api/shared/utils"
)

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
