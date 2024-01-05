package server

import (
	"exporia/internal/application/handler"
	"exporia/internal/domain/enum"
	"exporia/internal/middleware"
	"github.com/gin-gonic/gin"
)

type WebServer struct {
	profileServerHandler handler.ProfileServerHandler
	authentication       handler.AuthenticationServerHandler
	middleware           middleware.Middleware
}

func NewWebServer(
	profileServerHandler handler.ProfileServerHandler,
	authentication handler.AuthenticationServerHandler,
	middleware middleware.Middleware,
) WebServer {
	s := WebServer{
		profileServerHandler,
		authentication,
		middleware,
	}
	return s
}
func (s *WebServer) SetupRoot() {
	router := gin.Default()

	router.POST("/login", s.authentication.Login)
	router.POST("/user/add", s.profileServerHandler.Create)
	router.POST("/activation", s.profileServerHandler.ActivateUser)

	user := router.Group("/profil", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleUser, enum.RoleManager, enum.RoleAdmin}))
	user.PUT("/", s.profileServerHandler.Update)
	user.PUT("/pass/", s.profileServerHandler.UpdatePassword)
	user.DELETE("/", s.profileServerHandler.Delete)
	user.GET("/", s.profileServerHandler.GetUser)

	changeUserRole := router.Group("/rol", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleUser}))
	changeUserRole.GET("/", s.keyServerHandler.UpdateUserRole)

	router.Run("0.0.0.0:8001")
}
