package usecase

import (
	"context"
	"errors"
	"github.com/dungnh3/bpp-resolve/internal/domain/model"
	"github.com/dungnh3/bpp-resolve/internal/domain/repository"
	"github.com/dungnh3/bpp-resolve/internal/dto"
	"github.com/go-logr/logr"
	"github.com/shopspring/decimal"
)

var (
	ErrRequestInvalid = errors.New("buying_price must be lesser or equal to current_selling_price of the wager_id")
)

type UseCase struct {
	logger logr.Logger
	repo   repository.IRepository
}

func NewUseCase(logger logr.Logger, repo repository.IRepository) *UseCase {
	return &UseCase{
		logger: logger,
		repo:   repo,
	}
}

func (u *UseCase) InitializeWager(ctx context.Context, wagerDto *dto.CreateWagerDto) (*model.Wager, error) {
	wager := &model.Wager{
		TotalWagerValue:     wagerDto.TotalWagerValue,
		Odds:                wagerDto.Odds,
		SellingPercentage:   wagerDto.SellingPercentage,
		SellingPrice:        wagerDto.SellingPrice,
		CurrentSellingPrice: wagerDto.SellingPrice,
	}
	if err := u.repo.InitializeWager(ctx, wager); err != nil {
		return nil, err
	}
	return wager, nil
}

func (u *UseCase) FindWager(ctx context.Context, offset, limit int) ([]*model.Wager, error) {
	return u.repo.FindWagers(ctx, offset, limit)
}

func (u *UseCase) BuyWager(ctx context.Context, wagerID uint32, buyingPrice decimal.Decimal) (*model.Purchase, error) {

	purchase := &model.Purchase{
		WagerID:     wagerID,
		BuyingPrice: buyingPrice,
	}

	if err := u.repo.Transaction(func(r repository.IRepository) error {
		wager, err := r.SelectForUpdateWagerByID(ctx, wagerID)
		if err != nil {
			return err
		}

		// buying_price must be lesser or equal to current_selling_price of the wager_id
		if buyingPrice.GreaterThan(wager.CurrentSellingPrice) {
			return ErrRequestInvalid
		}

		currentSellingPrice := wager.CurrentSellingPrice.Sub(buyingPrice)

		amountSold := buyingPrice
		if wager.AmountSold != nil {
			amountSold = amountSold.Add(*wager.AmountSold)
		}

		percentageSold := amountSold.Mul(decimal.NewFromInt(100)).Div(wager.SellingPrice)

		if err := r.RecordPurchasingWagerByID(ctx, wagerID, currentSellingPrice, amountSold, percentageSold); err != nil {
			return err
		}

		if err := r.RecordPurchase(ctx, purchase); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return purchase, nil
}
