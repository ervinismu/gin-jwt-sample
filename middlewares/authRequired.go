package middlewares

import (
	"net/http"

	authjwt "github.com/ervinismu/gin-jwt-sample/helpers/auth_jwt"
	"github.com/ervinismu/gin-jwt-sample/initializers"
	"github.com/ervinismu/gin-jwt-sample/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		token, err := authjwt.Decode(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "error": "invalid token" })
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var user models.User
			initializers.DB.Find(&user, claims["user_id"])

			if user.ID == 0 {
				log.Error(err.Error())
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{ "error": "invalid token" })
			}

			// attach to req
			ctx.Set("user", user)

			// continue
			ctx.Next()
		} else {
			log.Error(err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "error": "invalid token" })
			return
		}
	}
}
