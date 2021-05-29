package usecase

import (
	"context"
	"github.com/dungnh3/bpp-resolve/internal/domain/model"
	"github.com/dungnh3/bpp-resolve/internal/dto"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *TestServiceSuite) TestUseCase_InitializeWager() {
	type args struct {
		ctx      context.Context
		wagerDto *dto.CreateWagerDto
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Wager
		wantErr error
	}{
		{
			name: "Testing => InitializeWager() success",
			args: args{
				ctx: context.Background(),
				wagerDto: &dto.CreateWagerDto{
					TotalWagerValue:   1000,
					Odds:              20,
					SellingPercentage: 50,
					SellingPrice:      decimal.NewFromInt(800),
				},
			},
			want: &model.Wager{
				TotalWagerValue:     1000,
				Odds:                20,
				SellingPercentage:   50,
				SellingPrice:        decimal.NewFromInt(800),
				CurrentSellingPrice: decimal.NewFromInt(800),
				PercentageSold:      nil,
				AmountSold:          nil,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got, err := suite.svc.InitializeWager(tt.args.ctx, tt.args.wagerDto)
			if err != tt.wantErr {
				suite.Fail("InitializeWager() failed", "error", err, "wantErr", tt.wantErr)
				return
			}

			if got != nil {
				assert.Equal(suite.T(), tt.want.TotalWagerValue, got.TotalWagerValue)
				assert.Equal(suite.T(), tt.want.Odds, got.Odds)
				assert.Equal(suite.T(), tt.want.SellingPercentage, got.SellingPercentage)
				assert.Equal(suite.T(), tt.want.SellingPrice, got.SellingPrice)
				assert.Equal(suite.T(), tt.want.CurrentSellingPrice, got.CurrentSellingPrice)
			}
		})
	}
}

func (suite *TestServiceSuite) TestUseCase_BuyWager() {
	type args struct {
		ctx         context.Context
		wager       *dto.CreateWagerDto
		buyingPrice decimal.Decimal
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Testing => BuyWager() success",
			args: args{
				ctx: context.Background(),
				wager: &dto.CreateWagerDto{
					TotalWagerValue:   1000,
					Odds:              40,
					SellingPercentage: 50,
					SellingPrice:      decimal.NewFromInt(600),
				},
				buyingPrice: decimal.NewFromInt(100),
			},
			wantErr: nil,
		}, {
			name: "Testing => BuyWager() failed",
			args: args{
				ctx: context.Background(),
				wager: &dto.CreateWagerDto{
					TotalWagerValue:   1000,
					Odds:              40,
					SellingPercentage: 50,
					SellingPrice:      decimal.NewFromInt(600),
				},
				buyingPrice: decimal.NewFromInt(610),
			},
			wantErr: ErrRequestInvalid,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			wager, err := suite.svc.InitializeWager(tt.args.ctx, tt.args.wager)
			require.NoError(suite.T(), err)

			purchase, err := suite.svc.BuyWager(tt.args.ctx, wager.ID, tt.args.buyingPrice)
			if err != tt.wantErr {
				suite.Fail("BuyWager() failed", "error", err, "wantErr", tt.wantErr)
				return
			}
			if purchase != nil {
				assert.Equal(suite.T(), wager.ID, purchase.WagerID)
				assert.Equal(suite.T(), tt.args.buyingPrice, purchase.BuyingPrice)
			}
		})
	}
}

func (suite *TestServiceSuite) TestUseCase_FindWager() {
	type args struct {
		ctx    context.Context
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "Testing => FindWager()",
			args: args{
				ctx:    context.Background(),
				offset: 0,
				limit:  1,
			},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got, err := suite.svc.FindWager(tt.args.ctx, tt.args.offset, tt.args.limit)
			if err != tt.wantErr {
				suite.Fail("FindWager() failed", "error", err, "wantErr", tt.wantErr)
				return
			}
			assert.Equal(suite.T(), tt.want, len(got))
		})
	}
}
