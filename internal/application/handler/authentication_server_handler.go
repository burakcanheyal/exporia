package handler

import (
	"exporia/internal/domain/dto"
	"exporia/internal/domain/service"
	"exporia/platform/validation"
	"exporia/platform/zap"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenticationServerHandler struct {
	authenticationService service.Authentication
}

func NewAuthenticationServerHandler(authenticationService service.Authentication) AuthenticationServerHandler {
	a := AuthenticationServerHandler{authenticationService}
	return a
}

func (u *AuthenticationServerHandler) Login(context *gin.Context) {
	user := dto.AuthDto{}
	if err := context.BindJSON(&user); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	tokens, err := u.authenticationService.Login(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	zap.Logger.Info("Giriş başarılı")
	context.JSON(http.StatusOK, tokens)
}
