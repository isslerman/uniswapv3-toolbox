package position

import (
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
}

func TestGetAmount0(t *testing.T) {
	for _, pos := range positions {
		amount0, price, minPrice, maxPrice := pos.amount0, pos.price, pos.minPrice, pos.maxPrice
		amount1 := GetAmount1(amount0, price, minPrice, maxPrice)
		assert.Equal(t, pos.amount1, amount1)
	}
}

func TestGetAmount1(t *testing.T) {
	for _, pos := range positions {
		amount1, price, minPrice, maxPrice := pos.amount1, pos.price, pos.minPrice, pos.maxPrice
		amount0 := GetAmount0(amount1, price, minPrice, maxPrice)
		assert.Equal(t, pos.amount0, amount0)
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
	amounts0 := 2.0
	amounts1 := 4000.00
	maxPrice := 3000.00
	price := 2000.00
	minPrice := GetMinPrice(amounts0, amounts1, price, maxPrice)
	assert.Equal(t, 1333.3333333333333, minPrice)
}
