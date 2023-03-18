package main

import (
	"net/http"
	"os"

	"github.com/ervinismu/gin-jwt-sample/controllers"
	"github.com/ervinismu/gin-jwt-sample/initializers"
	"github.com/ervinismu/gin-jwt-sample/middlewares"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()

	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})

	r.Use(middlewares.LoggingMiddleware())

	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Signin)

	r.Use(middlewares.AuthMiddleware())

	r.GET("/me", controllers.MyProfile)
	r.Run()
}
