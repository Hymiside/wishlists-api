package handler

import (
	"errors"
	"fmt"
	"github.com/Hymiside/wishlists-api/pkg/service"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	ErrUserIdNotFound = errors.New("userId not found")
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
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s]	REQUEST: %s %s    STATUS-CODE: %d    LATENSY: %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/profile", h.profile)
		api.GET("/favorites")
		api.GET("/user/:nickname")

	}

	return router
}
