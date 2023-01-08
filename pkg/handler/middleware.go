package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	schoolCtx           = "schoolId"
	authorizationHeader = "Authorization"
)

// schoolIdentity инденцифицирует пользователя при запросах в пути /api/...
func (h *Handler) schoolIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		responseWithError(c, http.StatusUnauthorized, "empty auth header")
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		responseWithError(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	if len(headerParts[1]) == 0 {
		responseWithError(c, http.StatusUnauthorized, "token is empty")
		return
	}
	schoolId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		responseWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(schoolCtx, schoolId)
}

func getSchoolId(c *gin.Context) (string, error) {
	data, ok := c.Get(schoolCtx)
	if !ok {
		return "", errors.New("schoolId not found")
	}
	schoolId := data.(string)
	return schoolId, nil
}
