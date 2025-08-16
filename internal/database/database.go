// internal/db/models.go
package database

import (
	"homelab-dashboard/internal/config"
	"homelab-dashboard/internal/logger"
	"homelab-dashboard/internal/utils"
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

	// Populate Admin User credentials
	var adminRole Role
	result := DB.Where("id = 1").Find(&adminRole)
	if result.Error != nil {
		logger.Log.Fatal(result.Error)
	}
	if result.RowsAffected != 0 {
		return
	}	

	adminRole = Role{
		Id: 1,
		Name: "admin",
	}
	result = DB.Create(&adminRole)
	if result.Error != nil {
		logger.Log.Fatal(result.Error)
		return
	}

	adminUser := User{
		Id: 1,
		Username: config.App.Server.AdminAuth.Username,
		PasswordHash: utils.HashPassword(config.App.Server.AdminAuth.Password),
	}
	result = DB.Create(&adminUser)
	if result.Error != nil {
		logger.Log.Fatal(result.Error)
	}

	adminUserRole := UserRole{
		RoleId: 1,
		UserId: 1,
	}
	result = DB.Create(&adminUserRole)
	if result.Error != nil {
		logger.Log.Fatal(result.Error)
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

	UserRoles      []UserRole       `gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE"`
	LoginSessions  []LoginSessions  `gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE"`
}

type Role struct {
    Id             uint       `gorm:"primarykey"`
    Name           RoleName   `gorm:"index"`
    CreatedAt      time.Time 

	UserRoles      []UserRole `gorm:"foreignKey:RoleId;references:Id;constraint:OnDelete:CASCADE"`
}

type UserRole struct {
	RoleId     uint 
	UserId     uint
    CreatedAt  time.Time

	Role       Role      `gorm:"foreignKey:Id;references:RoleId"`
	User       User      `gorm:"foreignKey:Id;references:UserId"`
}

type LoginSessions struct {
	UserId     uint 
	IP         string
    CreatedAt  time.Time

	User       User      `gorm:"foreignKey:Id;references:UserId"`
}
