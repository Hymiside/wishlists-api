package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		responseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseSuccessful(c, profile)
}
