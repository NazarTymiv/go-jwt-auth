package initializers

import "github.com/nazartymiv/jwt-auth/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}