package usecase

import (
	"context"
	"errors"
	"github.com/dungnh3/bpp-resolve/internal/domain/model"
	"github.com/dungnh3/bpp-resolve/internal/domain/repository"
	"github.com/dungnh3/bpp-resolve/internal/dto"
	"github.com/go-logr/logr"
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

func (u *UseCase) BuyWager(ctx context.Context, wagerID uint32, buyingPrice float64) (*model.Purchase, error) {
	wager, err := u.repo.FindWagerByID(ctx, wagerID)
	if err != nil {
		return nil, err
	}

	// buying_price must be lesser or equal to current_selling_price of the wager_id
	if buyingPrice > wager.CurrentSellingPrice {
		return nil, errors.New("request invalid! buying_price must be lesser or equal to current_selling_price")
	}

	purchase := &model.Purchase{
		WagerID:     wagerID,
		BuyingPrice: buyingPrice,
	}

	if err = u.repo.Transaction(func(r repository.IRepository) error {
		if err := r.RecordWagerPriceByID(ctx, wagerID, buyingPrice, wager.SellingPrice); err != nil {
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
