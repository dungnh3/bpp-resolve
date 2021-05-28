package server

import (
	"github.com/dungnh3/bpp-resolve/internal/service"
	"github.com/gin-gonic/gin"
)

func (s *Server) Register(svc *service.Service) error {
	gin.SetMode(gin.ReleaseMode)
	s.router = gin.Default()

	healthGr := s.router.Group("/health")
	{
		healthGr.GET("/ready", svc.Readiness)
		healthGr.GET("/lively", svc.Liveness)
	}

	wagerGr := s.router.Group("/wagers")
	{
		wagerGr.POST("", svc.InitializeWager)
		wagerGr.GET("", svc.ListWagers)
	}

	buyGr := s.router.Group("/buy/:wager_id")
	{
		buyGr.POST("", svc.BuyWager)
	}
	return nil
}
