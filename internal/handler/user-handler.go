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

func (ush *UserHandler) ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/profile/profile.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (ush *UserHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
	tmpl, err := template.ParseFiles("web/auth/register.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (ush *UserHandler) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/auth/login.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (ush *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Internal error", http.StatusBadGateway)
		return
	}

	status_code := ush.authService.Register(user)
	ResponseRegistration(status_code, w, r)
}

func (ush *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")

	var user models.User
	error_decode := json.NewDecoder(r.Body).Decode(&user)
	if error_decode != nil {
		http.Error(w, "Internal error", http.StatusBadGateway)
		return
	}

	status_code, id := ush.authService.Login(user)
	if status_code == 200 {
		jwt_token, err := ush.authService.Generate_Token(id)
		if err != nil {
			http.Error(w, "Internal error", http.StatusBadGateway)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt-token",
			Value:    jwt_token,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   1600 * 20,
		})
		ResponseLogin(status_code, w, r)
	} else {
		ResponseLogin(status_code, w, r)
	}

}
