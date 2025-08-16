package server

import (
	"encoding/json"
	"fmt"
	"homelab-dashboard/internal/logger"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func WriteJson(w http.ResponseWriter, j any) {
	jsonBytes, err := json.Marshal(j)
	if err != nil {
		logger.Log.Error(err)
	}
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(jsonBytes))
}

func ParseJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, JWT_SECRET_FUNC, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims); 
	if !ok {
		return nil, fmt.Errorf("Invalid claims")
	}

	return claims, nil
}

