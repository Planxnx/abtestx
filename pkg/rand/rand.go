// nolint: gocritic,gomnd
package rand

import (
	"crypto/rand"
	"math"
	"math/big"
)

func Intn(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

func Float64() float64 {
again:
	f := float64(Intn(math.MaxInt64)) / (1 << 63)
	if f == 1 {
		goto again
	}
	return f
}

func Floats64n(max float64, min ...float64) float64 {
	m := 0.0
	if len(min) > 0 {
		m = min[0]
	}
	return m + Float64()*(max-m)
}
