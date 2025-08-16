package server

import (
	"encoding/json"
	"homelab-dashboard/internal/database"
	"homelab-dashboard/internal/logger"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//------------------ Authentication Handler ------------------

type AuthAPIHandlers struct{} 

func (ah *AuthAPIHandlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(400)
		WriteJson(w, JSON{"error": "Invalid data"})
		return
	}

	username := data["username"]
	password := data["password"]

	var user database.User
	result := database.DB.Preload("UserRoles.Role").Where("username = ?", username).Limit(1).Find(&user)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			w.WriteHeader(401)
			WriteJson(w, JSON{"error": "Invalid credentials"})
			return
		}

		logger.Log.Error(result.Error)
		w.WriteHeader(500)
		WriteJson(w, JSON{"error": "Something went wrong"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
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

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: signedToken,
		Expires: expiresAt,
		HttpOnly: true,
		Path: "/",
	})
	w.WriteHeader(200)
	WriteJson(w, JSON{"success": true})
}

