package handler

import (
	"github.com/Nurt0re/chatik/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/Nurt0re/chatik/pkg/ws"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(wsHandler *ws.Handler) *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
		auth.POST("oauth", h.oAuth)
		auth.POST("callback", h.callback)
		auth.GET("logout", h.Logout)
	}
	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUser)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}
	}

	ws:= router.Group("/ws", h.userIdentity)
	{
		ws.POST("/createRoom", wsHandler.CreateRoom)
		ws.GET("/joinRoom/:roomId", wsHandler.JoinRoom)
	}
	return router
}
