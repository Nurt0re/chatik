package handler

import (
	"github.com/Nurt0re/chatik/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
		auth.POST("oauth", h.oAuth)
		auth.POST("callback", h.callback)
	}
	api := router.Group("/api", h.userIdentity)
	{
		chats := api.Group("/chats")
		{
			chats.POST("/")
		}

		users := api.Group("/users")
		{
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUser)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}
	}
	return router
}
