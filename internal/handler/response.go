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
	switch {
	case status_code == 200:
		FormatIntoJson("Account has been created!", w)
	case status_code == 400:
		FormatIntoJson("Account with this username is already exist!", w)
	case status_code == 406:
		FormatIntoJson("Username or Password is too short!(6 symbols for password, username len > 3)", w)
	case status_code == 502:
		FormatIntoJson("Internal error!", w)
	}
}

func ResponseLogin(status_code int, w http.ResponseWriter, r *http.Request) {
	switch {
	case status_code == 200:
		FormatIntoJson("You has been successfully logined!", w)
	case status_code == 400:
		FormatIntoJson("User with this username doesn't exist!", w)
	case status_code == 406:
		FormatIntoJson("Invalid password!", w)
	case status_code == 404:
		FormatIntoJson("Invalid username or password!", w)
	}
}

func ResponseArticle(status_code int, w http.ResponseWriter, r *http.Request) {
	switch {
	case status_code == 200:
		FormatIntoJson("Article has been created!", w)
	case status_code == 400:
		FormatIntoJson("Article's title or content too short!", w)
	case status_code == 403:
		FormatIntoJson("You don't have permission to delete/edit this article!", w)
	}
}

func ResponseUsernameChange(status_code int, w http.ResponseWriter, r *http.Request) {
	switch {
		case status_code == 200:
			FormatIntoJson("Your username has been updated!", w)
		case status_code == 400:
			FormatIntoJson("Username is too short!", w)
		case status_code == 409:
			FormatIntoJson("Username is already in use!", w)
	}
}

func ResponseBioChange(status_code int, w http.ResponseWriter, r *http.Request) {
	switch {
		case status_code == 200:
			FormatIntoJson("Your bio has been updated!", w)
		case status_code == 502:
			FormatIntoJson("Server error!", w)
	}
}

func ResponsePasswordChange(status_code int, w http.ResponseWriter, r *http.Request) {
	switch {
		case status_code == 200:
			FormatIntoJson("Password has been updated!", w)
		case status_code == 400:
			FormatIntoJson("Incorrect password!", w)
	}
}
