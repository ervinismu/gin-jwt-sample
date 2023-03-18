package controllers

import (
	"net/http"
	"time"

	authjwt "github.com/ervinismu/gin-jwt-sample/helpers/auth_jwt"
	reason "github.com/ervinismu/gin-jwt-sample/helpers/reason"
	"github.com/ervinismu/gin-jwt-sample/initializers"
	"github.com/ervinismu/gin-jwt-sample/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Email    string
	Password string
}

func Signup(ctx *gin.Context) {
	var req SignupRequest
	if ctx.Bind(&req) != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{ "error": reason.ErrFailedRegisterUser.Error()})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{ "error": reason.ErrFailedRegisterUser.Error()})

		return
	}

	user := models.User{Email: req.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{ "error": reason.ErrFailedRegisterUser.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func Signin(ctx *gin.Context) {
	var req SignupRequest
	if ctx.Bind(&req) != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{ "error": reason.ErrFailedLoginUser.Error() })

		return
	}

	var user models.User
	initializers.DB.Find(&user, "email = ?", req.Email)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{ "error": reason.ErrFailedLoginUser.Error() })

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{ "error": reason.ErrFailedLoginUser.Error() })

		return
	}

	exp := time.Now().Add(time.Minute * 30).Unix()
	userID := user.ID
	token, err := authjwt.Encode(exp, userID)
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{ "error": reason.ErrFailedLoginUser.Error() })

		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "token": token } )
}
