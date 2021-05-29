package service

import "errors"

type Error error

var (
	// ErrBadRequest
	ErrBadRequest = errors.New("bad request")
	// ErrRequestInValid
	ErrRequestInValid = errors.New("request invalid")

	// ErrTotalWagerValueMustGreaterThan0
	ErrTotalWagerValueMustGreaterThan0 = errors.New("total_wager_value must be specified as a positive integer above 0")
	// ErrOddsValueMustGreaterThan0
	ErrOddsValueMustGreaterThan0 = errors.New("odds must be specified as a positive integer above 0")
	// ErrSellingPercentageValue
	ErrSellingPercentageValue = errors.New("selling_percentage must be specified as an integer between 1 and 100")
	// ErrSellingPriceValueMustGreaterThan0
	ErrSellingPriceValueMustGreaterThan0 = errors.New("must be specified as a positive decimal value to two decimal places")
	// ErrSellingPriceValue
	ErrSellingPriceValue = errors.New("must be greater than total_wager_value * (selling_percentage / 100)")
)
