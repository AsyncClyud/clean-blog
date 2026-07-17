package handler

import (
	"encoding/json"
	"net/http"
)

func FormatIntoJson(message string, w http.ResponseWriter) {
	response, err := json.MarshalIndent(message, "", " ")
	if err != nil {
		http.Error(w, "Internal error", http.StatusBadGateway)
		return
	}

	error_encode := json.NewEncoder(w).Encode(string(response))
	if error_encode != nil {
		http.Error(w, "Internal error", http.StatusBadGateway)
		return
	}
}

func ResponseRegistration(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson("Account has been created!", w)
	case http.StatusBadRequest:
		FormatIntoJson("Username must be at least 4 characters long!", w)
	case http.StatusForbidden:
		FormatIntoJson("Captcha verification failed!", w)
	case http.StatusNotAcceptable:
		FormatIntoJson("Username can only contain letters, numbers!", w)
	case http.StatusConflict:
		FormatIntoJson("Account with this username already exist!", w)
	case http.StatusUnprocessableEntity:
		FormatIntoJson("Password must be at least 6 characters long!", w)
	case http.StatusBadGateway:
		FormatIntoJson("Internal error!", w)
	}
}

func ResponseLogin(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson("You has been successfully logined!", w)
	case http.StatusBadRequest:
		FormatIntoJson("User with this username doesn't exist!", w)
	case http.StatusForbidden:
		FormatIntoJson("Captcha verification failed!", w)
	case http.StatusNotFound:
		FormatIntoJson("Invalid username or password!", w)
	case http.StatusNotAcceptable:
		FormatIntoJson("Username can only contain letters, numbers!", w)
	case http.StatusUnprocessableEntity:
		FormatIntoJson("Password must be at least 6 characters long!", w)
	}
}

func ResponseArticle(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson("Success!", w)
	case http.StatusBadRequest:
		FormatIntoJson("Article's title is too short!", w)
	case http.StatusUnauthorized:
		FormatIntoJson("You don't have permission to delete/edit this article!", w)
	case http.StatusForbidden:
		FormatIntoJson("Captcha verification required!", w)
	}
}

func ResponseComment(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson("Success!", w)
	case http.StatusBadRequest:
		FormatIntoJson("Comment content cannot be null!", w)
	case http.StatusForbidden:
		FormatIntoJson("Captcha verification required!", w)
	}
}

func ResponseUsernameChange(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson("Your username has been updated!", w)
	case http.StatusBadRequest:
		FormatIntoJson("Username is too short!", w)
	case http.StatusConflict:
		FormatIntoJson("Username is already in use!", w)
	}
}

func ResponseBioChange(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson("Your bio has been updated!", w)
	case http.StatusBadRequest:
		FormatIntoJson("Bio is too long! 2000 chars max!", w)
	}
}

func ResponsePasswordChange(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson("Password has been updated!", w)
	case http.StatusBadRequest:
		FormatIntoJson("Incorrect password!", w)
	}
}
