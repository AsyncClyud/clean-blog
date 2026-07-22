package userhandler

import (
	"blog/internal/models"
	"encoding/json"
	"net/http"
)

func FormatIntoJson(w http.ResponseWriter, status int, payload string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var message models.Message = models.Message{Message: payload}

	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		http.Error(w, "Internal error!", http.StatusInternalServerError)
	}

}

func ResponseRegistration(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(w, http.StatusOK, "Account has been created!")
	case http.StatusBadRequest:
		FormatIntoJson(w, http.StatusBadRequest, "Username must be at least 4 characters long!")
	case http.StatusForbidden:
		FormatIntoJson(w, http.StatusForbidden,"Captcha verification failed!")
	case http.StatusNotAcceptable:
		FormatIntoJson(w, http.StatusNotAcceptable, "Username can only contain letters, numbers!")
	case http.StatusConflict:
		FormatIntoJson(w, http.StatusConflict, "Account with this username already exist!")
	case http.StatusUnprocessableEntity:
		FormatIntoJson(w, http.StatusUnprocessableEntity, "Password must be at least 6 characters long!")
	case http.StatusBadGateway:
		FormatIntoJson(w, http.StatusInternalServerError, "Internal error!")
	}
}

func ResponseLogin(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(w, http.StatusOK, "You has been successfully logined!")
	case http.StatusBadRequest:
		FormatIntoJson(w, http.StatusBadRequest, "User with this username doesn't exist!")
	case http.StatusForbidden:
		FormatIntoJson(w, http.StatusForbidden,"Captcha verification failed!")
	case http.StatusNotFound:
		FormatIntoJson(w, http.StatusNotFound, "Invalid password!")
	case http.StatusNotAcceptable:
		FormatIntoJson(w, http.StatusNotAcceptable, "Username can only contain letters, numbers!")
	case http.StatusConflict:
		FormatIntoJson(w, http.StatusConflict, "Account with this username already exist!")
	case http.StatusUnprocessableEntity:
		FormatIntoJson(w, http.StatusUnprocessableEntity, "Password must be at least 6 characters long!")
	case http.StatusBadGateway:
		FormatIntoJson(w, http.StatusInternalServerError, "Internal error!")
	}
}

func ResponseUsernameChange(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(w, http.StatusOK, "Your username has been updated!")
	case http.StatusBadRequest:
		FormatIntoJson(w, http.StatusBadRequest, "Username is too short!")
	case http.StatusConflict:
		FormatIntoJson(w, http.StatusConflict, "Username already in use!")
	}
}

func ResponseBioChange(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(w, http.StatusOK, "Your bio has been updated!")
	case http.StatusBadRequest:
		FormatIntoJson(w, http.StatusBadRequest, "Bio is too long! 2000 chars max!")
	}
}

func ResponsePasswordChange(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(w, http.StatusOK, "Password has been updated!")
	case http.StatusBadRequest:
		FormatIntoJson(w, http.StatusBadRequest, "Incorrect password!")
	}
}
