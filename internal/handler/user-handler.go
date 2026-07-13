package handler

import (
	"blog/internal/contextutil"
	"blog/internal/models"
	"blog/internal/service"
	"encoding/json"
	"html/template"
	"net/http"
)

type UserHandler struct {
	authService service.AuthService
}

func NewUserHandler(service service.AuthService) *UserHandler {
	return &UserHandler{authService: service}
}

func (ush UserHandler) SecurityHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
}

func (ush *UserHandler) ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	ush.SecurityHeaders(w)
	tmpl, err := template.ParseFiles("web/profile/profile.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (ush *UserHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ush.SecurityHeaders(w)

	userID, ok := r.Context().Value(contextutil.UserIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, status_code := ush.authService.FetchUser(userID)
	if status_code != http.StatusOK {
		http.Error(w, "Internal error", http.StatusBadGateway)
		return
	}

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Internal error", http.StatusBadGateway)
		return
	}

}

func (ush *UserHandler) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	ush.SecurityHeaders(w)
	tmpl, err := template.ParseFiles("web/auth/register.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (ush *UserHandler) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	ush.SecurityHeaders(w)
	tmpl, err := template.ParseFiles("web/auth/login.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (ush *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")
	ush.SecurityHeaders(w)

	var user models.User
	error_decode := json.NewDecoder(r.Body).Decode(&user)
	if error_decode != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	status_code, id := ush.authService.Register(user)
	if status_code == 200 {
		ush.authService.SetTokenInCookie(w, id)
		ResponseRegistration(status_code, w, r)
	} else {
		ResponseRegistration(status_code, w, r)
	}
}

func (ush *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")
	ush.SecurityHeaders(w)

	var user models.User
	error_decode := json.NewDecoder(r.Body).Decode(&user)
	if error_decode != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	status_code, id := ush.authService.Login(user)
	if status_code == 200 {
		ush.authService.SetTokenInCookie(w, id)
		ResponseLogin(status_code, w, r)
	} else {
		ResponseLogin(status_code, w, r)
	}

}

func (ush *UserHandler) GetArticleAuthorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")
	ush.SecurityHeaders(w)

	cookie, exist := r.Cookie("jwt-token")
	if exist != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	claims, validation_err := ush.authService.Validate_Token(cookie.Value)
	if validation_err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var author models.Article
	error_decode := json.NewDecoder(r.Body).Decode(&author)
	if error_decode != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	if claims != author.Author {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

}
