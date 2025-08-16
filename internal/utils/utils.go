package utils

import (
	"homelab-dashboard/internal/logger"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	// bcrypt.DefaultCost is usually good enough
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error(err)
	}
	return string(hash)
}
