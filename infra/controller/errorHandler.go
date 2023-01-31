package controller

import (
	"net/http"

	"github.com/juliocesarscheidt/go-orm-api/shared/utils"
)

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	message := err.Error()
	utils.Logger.Infof(message)
	if message == "Not Found" {
		SendNotFound(w, r, message)
		return
	} else if message == "Bad Request" {
		SendBadRequest(w, r, message)
		return
	} else if message == "Invalid Name" || message == "Invalid Email" || message == "Invalid Password" {
		SendBadRequest(w, r, message)
		return
	} else if message == "Internal Server Error" {
		SendInternalServerError(w, r, message)
		return
	}
	SendInternalServerError(w, r, message)
}
