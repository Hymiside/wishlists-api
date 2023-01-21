package handler

import (
	"errors"
	"github.com/Hymiside/wishlists-api/pkg/models"
	"github.com/Hymiside/wishlists-api/pkg/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) profile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var profile map[string]string
	profile, err = h.services.GetProfile(userId)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			responseWithError(c, http.StatusBadRequest, err.Error())
			return
		}
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responseSuccessful(c, profile)
}

func (h *Handler) wishes(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var wishes []models.Wish
	wishes, err = h.services.GetWishes(userId)
	if err != nil {
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responseSuccessful(c, wishes)
}

func (h *Handler) createWish(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var data models.Wish
	if err = c.BindJSON(&data); err != nil {
		responseWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	if data.Title == "" || strconv.Itoa(data.Price) == "" || data.Link == "" {
		responseWithError(c, http.StatusBadRequest, ErrInvalidRequest.Error())
		return
	}
	data.UserId = userId

	var wishId string
	wishId, err = h.services.CreateWish(data)
	if err != nil {
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responseSuccessful(c, wishId)
}

func (h *Handler) favorites(c *gin.Context) {
	//userId, err := getUserId(c)
	//if err != nil {
	//	responseWithError(c, http.StatusUnauthorized, err.Error())
	//	return
	//}
}
