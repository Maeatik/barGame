package handler

import (
	"github.com/gin-gonic/gin"
	"userBar/back/service"
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
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
	}
	api := router.Group("/api", h.userIdentity)
	{
		client := api.Group("/client")
		{
			client.GET("/me", h.Me)

			client.GET("/list", h.ClientList)

			client.PUT("/buy", h.Buy)

		}
		barman := api.Group("/barman")
		{
			barman.GET("/list", h.BarmanList)

			barman.POST("/create", h.Create)
		}
	}

	return router
}
