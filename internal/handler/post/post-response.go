package posthandler

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

func ResponseArticle(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(w, http.StatusOK, "Success!")
	case http.StatusBadRequest:
		FormatIntoJson(w, http.StatusBadRequest, "Article title is too short!")
	case http.StatusUnprocessableEntity:
		FormatIntoJson(w, http.StatusBadRequest, "Article content is too short!")
	case http.StatusUnauthorized:
		FormatIntoJson(w, http.StatusUnauthorized, "You don't have permission to delete/edit this article!")
	case http.StatusForbidden:
		FormatIntoJson(w, http.StatusForbidden, "Captcha verification required!")
	}
}

func ResponseComment(status_code int, w http.ResponseWriter, r *http.Request) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(w, http.StatusOK, "Success!")
	case http.StatusBadRequest:
		FormatIntoJson(w, http.StatusBadRequest, "Comment content cannot be null!")
	case http.StatusForbidden:
		FormatIntoJson(w, http.StatusForbidden, "Captcha verification required!")
	}
}
