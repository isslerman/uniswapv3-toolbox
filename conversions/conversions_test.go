package conversions

import (
	"testing"
)

func TestGetTickToSqrtPriceX96(t *testing.T) {
	// tick, _ := NewTick(204632)

	// GetTickToSqrtPriceX96(0)
	// // test tick 0
	// zeroTick, _ := NewTick(0)
	// r0, _ := GetTickToSqrtPriceX96(zeroTick.Index)
	// assert.Equal(t, r0, new(big.Int).Lsh(big.NewInt(1), 96), "returns the correct value for tick 0")

	// // test max tick
	// maxTick, _ := NewTick(MaxTick)
	// rmin, _ := GetTickToSqrtPriceX96(maxTick.Index)
	// assert.Equal(t, rmin, MaxSqrtRatio, "returns the correct value for max tick")

	// tick, _ := NewTick(204632)
	// r1, _ := GetTickToSqrtPriceX96(tick.Index)
	// fmt.Printf("New value is: %v and type is: %T\n", r1, r1)

	// r2, _ := new(big.Int).SetString("fffcb933bd6fad37aa2d162d1a594001", 16)
	// fmt.Printf("New value is: %v and type is: %T\n", r2, r2)

	// tick, _ = NewTick(204632)
	// sqrtX96, _ := GetTickToSqrtPriceX96(tick.Index)
	// // result, _ := new(big.Int).SetString("2198795518956857915306525730013184", 10)
	// result, _ := new(big.Int).SetString("2198795518959335454840797035509975", 10)
	// assert.Equal(t, result, sqrtX96)

	// resultTick := GetPriceToTick(5000, 12)
	// assert.Equal(t, resultTick, int64(85176))
}
