package handler

import (
	"github.com/Hymiside/wishlists-api/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var data models.User

	if err := c.BindJSON(&data); err != nil {
		responseWithError(c, http.StatusBadRequest, "error to parse json")
		return
	}
	if data.Name == "" || data.Nickname == "" || data.Email == "" || data.Password == "" {
		responseWithError(c, http.StatusBadRequest, "error invalid request")
		return
	}
	userId, err := h.services.CreateUser(data)
	if err != nil {
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseSuccessful(c, map[string]string{"userId": userId})
}

func (h *Handler) signIn(context *gin.Context) {}
