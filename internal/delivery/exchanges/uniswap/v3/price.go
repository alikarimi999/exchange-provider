package uniswapv3

import (
	"errors"
	"math/big"
)

var (
	ErrDifferentCurrencies = errors.New("different currencies")
)

type Price struct {
	*Fraction
	BaseCurrency  token     // input i.e. denominator
	QuoteCurrency token     // output i.e. numerator
	Scalar        *Fraction // used to adjust the raw fraction w/r/t the decimals of the {base,quote}Token
}

// Construct a price, either with the base and quote currency amount, or the args
func NewPrice(baseCurrency, quoteCurrency token, denominator, numerator *big.Int) *Price {
	return &Price{
		Fraction:      NewFraction(numerator, denominator),
		BaseCurrency:  baseCurrency,
		QuoteCurrency: quoteCurrency,
		Scalar: NewFraction(
			new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(baseCurrency.Decimals)), nil),
			new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(quoteCurrency.Decimals)), nil)),
	}
}

// Invert flips the price, switching the base and quote currency
func (p *Price) Invert() *Price {
	return NewPrice(p.QuoteCurrency, p.BaseCurrency, p.Numerator, p.Denominator)
}

// AdjustedForDecimals Get the value scaled by decimals for formatting
func (p *Price) AdjustedForDecimals() *Fraction {
	return p.Fraction.Multiply(p.Scalar)
}

func (p *Price) ToSignificant(significantDigits int32) string {
	return p.AdjustedForDecimals().ToSignificant(significantDigits)
}

func (p *Price) ToFixed(decimalPlaces int32) string {
	return p.AdjustedForDecimals().ToFixed(decimalPlaces)
}
