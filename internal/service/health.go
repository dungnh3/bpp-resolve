package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Readiness health checking
func (s *Service) Readiness(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
	return
}

// Liveness health checking
func (s *Service) Liveness(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
	return
}
