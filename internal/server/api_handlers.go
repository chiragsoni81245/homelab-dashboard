package server

import (
	"errors"
	"homelab-dashboard/internal/database"
	"homelab-dashboard/internal/logger"
	"homelab-dashboard/internal/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

//------------------ Authentication Handler ------------------

type AuthAPIHandlers struct{} 

func (ah *AuthAPIHandlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	passwordHash := utils.HashPassword(password)

	var user database.User
	result := database.DB.Where("username = ?", username).Limit(1).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			w.WriteHeader(401)
			WriteJson(w, JSON{"error": "Invalid credentials"})
			return
		}

		logger.Log.Error(result.Error)
		w.WriteHeader(500)
		WriteJson(w, JSON{"error": "Something went wrong"})
		return
	}

	logger.Log.Info(username, password)

	if user.PasswordHash != passwordHash {
		w.WriteHeader(401)
		WriteJson(w, JSON{"error": "Invalid credentials"})
		return
	}

	userRoles := []string{}
	for _, role := range user.UserRoles {
		userRoles = append(userRoles, string(role.Role.Name))
	}

	expiresAt := time.Now().Add(24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		Username: username,
		UserId: user.Id,
		Roles: userRoles, 
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})

	signedToken, err := token.SignedString(JWT_SECRET)
	if err != nil {
		logger.Log.Error(err)
		WriteJson(w, JSON{"error": "Something went wrong"})	
		return
	}

	w.WriteHeader(200)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: signedToken,
		Expires: expiresAt,
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	WriteJson(w, JSON{"success": true})
}

