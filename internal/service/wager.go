package service

import (
	"github.com/dungnh3/bpp-resolve/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
)

func (s *Service) IsValidRequest(wager *dto.CreateWagerDto) error {
	if wager.TotalWagerValue <= 0 {
		return ErrTotalWagerValueMustGreaterThan0
	}

	if wager.Odds <= 0 {
		return ErrOddsValueMustGreaterThan0
	}

	if wager.SellingPercentage > 100 || wager.SellingPercentage < 1 {
		return ErrSellingPercentageValue
	}

	if wager.SellingPrice.LessThan(decimal.NewFromInt(0)) {
		return ErrSellingPriceValueMustGreaterThan0
	}

	// selling_price must be greater than total_wager_value * (selling_percentage / 100)
	sp := float64(wager.TotalWagerValue) * float64(wager.SellingPercentage) / 100
	if wager.SellingPrice.LessThanOrEqual(decimal.NewFromFloat(sp)) {
		return ErrSellingPriceValue
	}
	return nil
}

func (s *Service) InitializeWager(ctx *gin.Context) {
	var wagerDto dto.CreateWagerDto
	if err := ctx.ShouldBindJSON(&wagerDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	if err := s.IsValidRequest(&wagerDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
