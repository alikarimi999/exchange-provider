package calculate

import "math/big"

// y = (sqrt(x(4ad³ + x (4a(d - x) - d )²)) + x (4a(d - x) - d ))/8ax
// commonPart = 4a(d - x) - d
// sqrt = sqrt(x * (4ad³ + x * commonPart²)
// y =   (sqrt + x * commonPart) / divider
func getY(x *big.Float, a *big.Float, d *big.Float) *big.Float {
	dmx := new(big.Float).Sub(d, x)
	commonPartBig := new(big.Float).Sub(new(big.Float).Mul(new(big.Float).Mul(big.NewFloat(4), a), dmx), d)
	dCubed := new(big.Float).Mul(d, new(big.Float).Mul(d, d))

	commonPartSquared := new(big.Float).Mul(commonPartBig, commonPartBig)
	firstPart := new(big.Float).Add(new(big.Float).Mul(big.NewFloat(4),
		new(big.Float).Mul(a, dCubed)), new(big.Float).Mul(x, commonPartSquared))

	sqrtBig := new(big.Float).Sqrt(new(big.Float).Mul(x, firstPart))
	dividerBig := new(big.Float).Mul(big.NewFloat(8), new(big.Float).Mul(a, x))
	result := new(big.Float).Quo(new(big.Float).Add(new(big.Float).Mul(commonPartBig, x),
		sqrtBig), dividerBig)

	return result
}
