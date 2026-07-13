package service

import (
	_ "blog/internal/config"
	"blog/internal/models"
	"blog/internal/storage"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo       storage.UserRepository
	secret_key []byte
}

func NewAuthService(userDB storage.UserRepository, secret_key []byte) *AuthService {
	return &AuthService{repo: userDB, secret_key: secret_key}
}

func (ur *AuthService) HashPassword(password string) (hashed_password string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (ur *AuthService) CheckPaswordHash(password string, hashedPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (ur *AuthService) Generate_Token(id int) (token string, err error) {
	claims := jwt.MapClaims{
		"UserID": id,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	new_token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return new_token.SignedString(ur.secret_key)

}

func (ur *AuthService) Validate_Token(tokenString string) (Id int, err error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return ur.secret_key, nil
	})
	if err != nil {
		return 0, fmt.Errorf("Parsing error: %v", err)
	}
	if !token.Valid {
		return 0, fmt.Errorf("Invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("Invalid claims")
	}
	userIDValue, exists := claims["UserID"]
	if !exists {
		return 0, fmt.Errorf("Invalid UserID")
	}
	userID := userIDValue.(float64)

	return int(userID), nil
}

func (ur *AuthService) SetTokenInCookie(w http.ResponseWriter, id int) {
	jwt_token, err := ur.Generate_Token(id)
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
}

func (ur *AuthService) ValidateUserData(user models.User) (status_code int) {
	if len(user.Username) <= 2 {
		return http.StatusBadRequest
	}
	if len(user.Password) <= 5 {
		return http.StatusNotAcceptable
	}
	return http.StatusOK
}

func (ur *AuthService) FetchUser(user_id int) (user string, status_code int) {
	user, err := ur.repo.GetUserInfo(user_id)
	if err != nil {
		return "", http.StatusBadRequest
	}
	return user, http.StatusOK
}

func (ur *AuthService) Register(user models.User) (status_code, id int) {
	if ur.ValidateUserData(user) == 400 {
		return http.StatusNotAcceptable, 0
	}
	if ur.ValidateUserData(user) == 406 {
		return http.StatusNotAcceptable, 0
	}
	hashed_password, err := ur.HashPassword(user.Password)
	if err != nil {
		return http.StatusBadGateway, 0
	}
	new_user := models.User{Username: user.Username, Password: hashed_password}

	id, message := ur.repo.CreateUser(new_user)
	if !message {
		return http.StatusBadRequest, 0
	}
	log.Println(id)
	return http.StatusOK, id
}

func (ur *AuthService) Login(user models.User) (status_code, id int) {
	if ur.ValidateUserData(user) == 400 {
		return http.StatusBadRequest, 0
	}
	if ur.ValidateUserData(user) == 406 {
		return http.StatusNotAcceptable, 0
	}
	id, hashed_password, ok := ur.repo.CheckIfUserExist(user)
	if !ok {
		return http.StatusNotFound, 0
	}
	if ur.CheckPaswordHash(user.Password, hashed_password) != nil {
		return http.StatusNotFound, 0
	}
	return http.StatusOK, id
}
