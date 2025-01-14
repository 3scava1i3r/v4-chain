package lib

import (
	"math"
	"math/big"
)

const (
	OneMillion        = uint32(1_000_000)
	TenThousand       = uint32(10_000)
	OneHundred        = uint32(100)
	MaxPriceChangePpm = uint32(10_000)
	// 10^6 quantums == 1 USD.
	QuoteCurrencyAtomicResolution = int32(-6)
	UsdcAssetId                   = uint32(0)

	ZeroUint64 = uint64(0)
)

// BigInt0 returns a `big.Int` that is set to 0.
func BigInt0() *big.Int {
	return big.NewInt(0)
}

// BigNegMaxUint64 returns a `big.Int` that is set to -math.MaxUint64.
func BigNegMaxUint64() *big.Int {
	return new(big.Int).Neg(
		new(big.Int).SetUint64(math.MaxUint64),
	)
}

// BigMaxInt32 returns a `big.Int` that represents `MaxInt32`.
func BigMaxInt32() *big.Int {
	return big.NewInt(math.MaxInt32)
}

// BigFloat0 returns a `big.Float` that is set to 0.
func BigFloat0() *big.Float {
	return big.NewFloat(0)
}

// BigFloatMaxUint64 returns a `big.Float` that is set to MaxUint64.
func BigFloatMaxUint64() *big.Float {
	return new(big.Float).SetUint64(math.MaxUint64)
}

// BigIntOneMillion returns a `big.Int` that is set to 1_000_000.
func BigIntOneMillion() *big.Int {
	return big.NewInt(1_000_000)
}

// BigIntOneTrillion returns a `big.Int` that is set to 1_000_000_000_000.
func BigIntOneTrillion() *big.Int {
	return big.NewInt(1_000_000_000_000)
}

// BigRatOneMillion returns a `big.Rat` that is set to 1_000_000.
func BigRatOneMillion() *big.Rat {
	return big.NewRat(1_000_000, 1)
}

// BigRat0 returns a `big.Rat` that is set to 0.
func BigRat0() *big.Rat {
	return big.NewRat(0, 1)
}

// BigRat1 returns a `big.Rat` that is set to 1.
func BigRat1() *big.Rat {
	return big.NewRat(1, 1)
}
