package usecase

import (
	"context"
	"errors"
	"github.com/dungnh3/bpp-resolve/internal/domain/model"
	"github.com/dungnh3/bpp-resolve/internal/domain/repository"
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

func (u *UseCase) CreateWager(ctx context.Context, wager *model.Wager) error {
	return u.repo.RecordWager(ctx, wager)
}

func (u *UseCase) FindWager(ctx context.Context, offset, limit int) ([]*model.Wager, error) {
	return u.repo.ListWager(ctx, offset, limit)
}

func (u *UseCase) BuyWager(ctx context.Context, wagerID uint32, buyingPrice float64) error {
	wager, err := u.repo.FindWagerByID(ctx, wagerID)
	if err != nil {
		return err
	}

	// buying_price must be lesser or equal to current_selling_price of the wager_id
	if buyingPrice > wager.CurrentSellingPrice {
		return errors.New("request invalid! buying_price must be lesser or equal to current_selling_pric")
	}


	return nil
}
