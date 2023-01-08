package handler

import (
	"errors"
	"net/http"

	"github.com/Hymiside/wishlists-api/pkg/models"
	"github.com/Hymiside/wishlists-api/pkg/repository"
	"github.com/Hymiside/wishlists-api/pkg/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var data models.User

	if err := c.BindJSON(&data); err != nil {
		responseWithError(c, http.StatusBadRequest, ErrParseJSON.Error())
		return
	}
	if data.Name == "" || data.Nickname == "" || data.Email == "" || data.Password == "" {
		responseWithError(c, http.StatusBadRequest, ErrInvalidRequest.Error())
		return
	}
	userId, err := h.services.CreateUser(data)
	if err != nil {
		if errors.Is(err, repository.ErrUniqueKeyViolation) {
			responseWithError(c, http.StatusBadRequest, err.Error())
			return
		}
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseSuccessful(c, userId)
}

func (h *Handler) signIn(c *gin.Context) {
	var data models.User

	if err := c.BindJSON(&data); err != nil {
		responseWithError(c, http.StatusBadRequest, ErrParseJSON.Error())
		return
	}

	if data.Email == "" || data.Password == "" {
		responseWithError(c, http.StatusBadRequest, ErrInvalidRequest.Error())
		return
	}

	token, err := h.services.GenerateToken(data.Email, data.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidPwd) || errors.Is(err, repository.ErrItemsNotFound) {
			responseWithError(c, http.StatusBadRequest, err.Error())
			return
		}
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseSuccessful(c, token)
}
