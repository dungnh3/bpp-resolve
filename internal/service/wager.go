package service

import (
	"github.com/dungnh3/bpp-resolve/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Service) IsValidRequest(wager *dto.CreateWagerDto) bool {
	if wager.TotalWagerValue <= 0 || // total_wager_value must be specified as a positive integer above 0
		wager.Odds <= 0 || // odds must be specified as a positive integer above 0
		wager.SellingPercentage > 100 || wager.SellingPercentage < 1 || // selling_percentage must be specified as an integer between 1 and 100
		wager.SellingPrice < 0 { // selling_price must be specified as a positive decimal value to two decimal places, it is a monetary value
		return false
	}

	// selling_price must be greater than total_wager_value * (selling_percentage / 100)
	sp := float64(wager.TotalWagerValue) * float64(wager.SellingPercentage) / 100
	if wager.SellingPrice <= sp {
		return false
	}
	return true
}

func (s *Service) InitializeWager(ctx *gin.Context) {
	var wagerDto dto.CreateWagerDto
	if err := ctx.ShouldBindJSON(&wagerDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	if !s.IsValidRequest(&wagerDto) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrRequestInValid.Error()})
		return
	}

	wager, err := s.uc.InitializeWager(ctx, &wagerDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, wager)
	return
}

func (s *Service) ListWagers(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	offset := page - 1
	wagers, err := s.uc.FindWager(ctx, offset, limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	ctx.JSON(http.StatusOK, wagers)
	return
}

func (s *Service) BuyWager(ctx *gin.Context) {
	wagerId, err := strconv.Atoi(ctx.Param("wager_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	var buyingWagerDto dto.BuyingWagerDto
	if err = ctx.ShouldBindJSON(&buyingWagerDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	purchase, err := s.uc.BuyWager(ctx, uint32(wagerId), buyingWagerDto.BuyingPrice)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, purchase)
	return
}
