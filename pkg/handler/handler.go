package handler

import (
	"errors"

	"github.com/Hymiside/wishlists-api/pkg/service"
	"github.com/gin-gonic/gin"
)

var (
	ErrParseJSON      = errors.New("error to parse json")
	ErrInvalidRequest = errors.New("error invalid request")
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
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	return router
}
