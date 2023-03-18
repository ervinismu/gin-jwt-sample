package initializers

import "github.com/ervinismu/gin-jwt-sample/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
