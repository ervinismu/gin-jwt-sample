package initializers

import (
	"github.com/ervinismu/gin-jwt-sample/models"
	log "github.com/sirupsen/logrus"
)

func SyncDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Error("Cannot migrate database")
	}
}
