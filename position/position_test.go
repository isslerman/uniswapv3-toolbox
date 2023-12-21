package position

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var positions = []struct {
	amount0    float64
	amount1    float64
	price      float64
	minPrice   float64
	maxPrice   float64
	liquidity0 float64
	liquidity1 float64
}{
	{1.00, 2907.729524805772, 2486.80, 1994.20, 2998.90, 557.9599554712883, 557.9599554712883},
	{2.00, 4000.02224746262, 2000.00, 1333.33, 3000.00, 487.41718030204123, 487.41718030204123},
	{0.4570, 384.13600301246356, 2249.07, 2219.456, 2330.9121, 1226.2603874027805, 1226.2603874027805},
	{0.2089, 500, 2244.64, 2024.3866, 2472.5666, 0.0, 0.0},
	{0.210171, 500, 2243.52, 2024.3866, 2472.5666, 0.0, 0.0},
	{3.4248, 500, 155.81, 141.40800000000002, 172.832, 0.0, 0.0},
}

func TestGetAmount1(t *testing.T) {
	for _, pos := range positions {
		amount0, price, minPrice, maxPrice := pos.amount0, pos.price, pos.minPrice, pos.maxPrice
		amount1 := GetAmount1(amount0, price, minPrice, maxPrice)
		perc := 1 - (pos.amount1 / amount1)
		assert.LessOrEqual(t, perc, 0.005)
	}
}

func TestGetAmount0(t *testing.T) {
	for _, pos := range positions {
		amount1, price, minPrice, maxPrice := pos.amount1, pos.price, pos.minPrice, pos.maxPrice
		amount0 := GetAmount0(amount1, price, minPrice, maxPrice)
		perc := 1 - (pos.amount0 / amount0)
		assert.LessOrEqual(t, perc, 0.005)
		// assert.Equal(t, pos.amount0, amount0)
	}
}

func TestGetLiquidity0(t *testing.T) {
	for _, pos := range positions {
		amount0, price, maxPrice := pos.amount0, pos.price, pos.maxPrice
		l0 := GetLiquidity0(amount0, price, maxPrice)
		assert.Equal(t, pos.liquidity0, l0)
	}
}

func TestGetLiquidity1(t *testing.T) {
	for _, pos := range positions {
		amount1, price, minPrice := pos.amount1, pos.price, pos.minPrice
		l1 := GetLiquidity1(amount1, minPrice, price)
		assert.Equal(t, pos.liquidity1, l1)
	}
}

func TestGetMinPrice(t *testing.T) {
	for _, pos := range positions {
		amount0, amounts1, price, maxPrice := pos.amount0, pos.amount1, pos.price, pos.maxPrice
		minPrice := math.Floor(GetMinPrice(amount0, amounts1, price, maxPrice)*1000) / 1000
		assert.Equal(t, pos.minPrice, minPrice)
	}
}

func TestLiquidity(t *testing.T) {
	// for _, pos := range positions {
	pos := positions[3]
	amount0, amount1, price, lowPrice, highPrice := pos.amount0, pos.amount1, pos.price, pos.minPrice, pos.maxPrice
	newPrice := 157.12
	newAmount0, newAmount1 := GetAmountsWithLiquidity(amount0, amount1, price, newPrice, lowPrice, highPrice)
	assert.Equal(t, 10, newAmount0)
	assert.Equal(t, 10, newAmount1)
	// }
}
