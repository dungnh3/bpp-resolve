package service

import (
	"github.com/dungnh3/bpp-resolve/internal/dto"
	"github.com/shopspring/decimal"
	"testing"
)

func TestService_IsValidRequest(t *testing.T) {
	type args struct {
		wager *dto.CreateWagerDto
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Testing => IsValidRequest() verify arguments to initialize wager success",
			args: args{
				wager: &dto.CreateWagerDto{
					TotalWagerValue:   1000,
					Odds:              20,
					SellingPercentage: 50,
					SellingPrice:      decimal.NewFromInt(800),
				},
			},
			want: nil,
		}, {
			name: "Testing => IsValidRequest() failed, total_wager_value must be specified as a positive integer above 0",
			args: args{
				wager: &dto.CreateWagerDto{
					TotalWagerValue:   0,
					Odds:              20,
					SellingPercentage: 50,
					SellingPrice:      decimal.NewFromInt(800),
				},
			},
			want: ErrTotalWagerValueMustGreaterThan0,
		}, {
			name: "Testing => IsValidRequest() failed, odds must be specified as a positive integer above 0",
			args: args{
				wager: &dto.CreateWagerDto{
					TotalWagerValue:   1000,
					Odds:              0,
					SellingPercentage: 50,
					SellingPrice:      decimal.NewFromInt(800),
				},
			},
			want: ErrOddsValueMustGreaterThan0,
		}, {
			name: "Testing => IsValidRequest() failed, selling_percentage must be specified as an integer between 1 and 100",
			args: args{
				wager: &dto.CreateWagerDto{
					TotalWagerValue:   1000,
					Odds:              20,
					SellingPercentage: 120,
					SellingPrice:      decimal.NewFromInt(800),
				},
			},
			want: ErrSellingPercentageValue,
		}, {
			name: "Testing => IsValidRequest() failed, selling_price must be specified as a positive decimal value to two decimal places, it is a monetary value",
			args: args{
				wager: &dto.CreateWagerDto{
					TotalWagerValue:   1000,
					Odds:              20,
					SellingPercentage: 50,
					SellingPrice:      decimal.NewFromInt(-100),
				},
			},
			want: ErrSellingPriceValueMustGreaterThan0,
		}, {
			name: "Testing => IsValidRequest() failed, selling_price must be greater than total_wager_value * (selling_percentage / 100)",
			args: args{
				wager: &dto.CreateWagerDto{
					TotalWagerValue:   1000,
					Odds:              20,
					SellingPercentage: 50,
					SellingPrice:      decimal.NewFromInt(500),
				},
			},
			want: ErrSellingPriceValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if got := s.IsValidRequest(tt.args.wager); got != tt.want {
				t.Errorf("IsValidRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
