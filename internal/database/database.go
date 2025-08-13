// internal/db/models.go
package database

import (
	"homelab-dashboard/internal/logger"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(path string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		logger.Log.Fatalf("failed to connect to database: %v", err)
	}

	for _, table := range []any{&User{}, &Role{}, &UserRole{}, &LoginSessions{}} {
		err = DB.AutoMigrate(table)
		if err != nil {
			logger.Log.Fatalf("failed to run migrations: %v", err)
		}
	}
}


type RoleName string
const (
    AdminRole RoleName = "admin"
)

type User struct {
    Id             uint             `gorm:"primarykey"`
    Username       string           `gorm:"index"`
    PasswordHash   string           `gorm:"index"`
    CreatedAt      time.Time

	UserRoles      []UserRole       `gorm:"foreignKey:UserId"`
	LoginSessions  []LoginSessions  `gorm:"foreignKey:UserId"`
}

type Role struct {
    Id             uint       `gorm:"primarykey"`
    Name           RoleName   `gorm:"index"`
    PasswordHash   string     `gorm:"index"`
    CreatedAt      time.Time 

	UserRoles      []UserRole `gorm:"foreignKey:RoleId"`
}

type UserRole struct {
	RoleId     uint 
	UserId     uint
    CreatedAt  time.Time

	Role       Role      `gorm:"foreignKey:Id"`
	User       User      `gorm:"foreignKey:Id"`
}

type LoginSessions struct {
	UserId     uint 
	IP         string
    CreatedAt  time.Time

	User       User      `gorm:"foreignKey:Id"`
}
