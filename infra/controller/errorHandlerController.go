package controller

import (
	"net/http"
	"strings"

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
	} else if message == "Invalid password length" {
		SendUnprocessableEntity(w, r, "Invalid password length, the password must have at least 8 and at most 50 characters")
		return
	} else if strings.Contains(message, "Invalid name") ||
		strings.Contains(message, "Invalid email") ||
		strings.Contains(message, "Invalid password") {
		SendBadRequest(w, r, message)
		return
	} else if message == "Internal Server Error" {
		SendInternalServerError(w, r, message)
		return
	}
	SendInternalServerError(w, r, message)
}
