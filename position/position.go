package position

import (
	"math"
)

// If you have an open position and the liquidity, calculate the amount0
func GetAmount0WithLiquidity(liquidity, price, minPrice, maxPrice float64) float64 {
	sPrice := math.Sqrt(price)
	sLowerRange := math.Sqrt(minPrice)
	sUpperRange := math.Sqrt(maxPrice)
	// if the price is outside the range, use range
	sPrice = math.Max(math.Min(sPrice, sUpperRange), sLowerRange)
	return liquidity * (sUpperRange - sPrice) / (sPrice * sUpperRange)
}

func GetAmount1WithLiquidity(liquidity, price, minPrice, maxPrice float64) float64 {
	sPrice := math.Sqrt(price)
	sLowerRange := math.Sqrt(minPrice)
	sUpperRange := math.Sqrt(maxPrice)
	// if the price is outside the range, use range
	sPrice = math.Max(math.Min(sPrice, sUpperRange), sLowerRange)
	return liquidity * (sPrice - sLowerRange)
}

// If you have an open position and want to know the new amount of the assets after the price move.
func GetAmountsWithLiquidity(amount0, amount1, price, newPrice, minPrice, maxPrice float64) (float64, float64) {
	L := GetLiquidity(amount0, amount1, price, minPrice, maxPrice)
	L0 := GetAmount0WithLiquidity(L, newPrice, minPrice, maxPrice)
	L1 := GetAmount1WithLiquidity(L, newPrice, minPrice, maxPrice)
	return L0, L1
}

// Get amount1 if you have the amount0, current price, min price and max price range
func GetAmount1(amount0 float64, price float64, minPrice float64, maxPrice float64) float64 {
	l0 := GetLiquidity0(amount0, price, maxPrice)
	return l0 * (math.Sqrt(price) - math.Sqrt(minPrice))
}

// Get amount0 if you have the amount1, current price, min price and max price range
func GetAmount0(amount1 float64, price float64, minPrice float64, maxPrice float64) float64 {
	l1 := GetLiquidity1(amount1, price, minPrice)
	return l1 * (math.Sqrt(maxPrice) - math.Sqrt(price)) / (math.Sqrt(price) * math.Sqrt(maxPrice))
}

// If you have an open position, calculaty the liquidity
func GetLiquidity(amounts0, amounts1, price, lowerRange, upperRange float64) float64 {
	if price <= lowerRange {
		return GetLiquidity0(amounts0, lowerRange, upperRange)
	} else if price < upperRange {
		liq0 := GetLiquidity0(amounts0, price, upperRange)
		liq1 := GetLiquidity1(amounts1, lowerRange, price)
		return math.Min(liq0, liq1)
	} else {
		return GetLiquidity1(amounts1, lowerRange, upperRange)
	}
}

// The goal is to have an optimally balanced position, where Liquidity0 == Liquidity1.
// (Explanation: to calculate the liquidity of a position where the current price is within the price range, Uniswap uses
// the minimum of the liquidities provided by the two tokens in that position. If the amount of one token is more than necessary,
// the extra liquidity provided is essentially ignored from a LP perspective. So your goal to have such an amount of y in the pool
// such that the liquidity of y exactly matches the liquidity of x.
func GetLiquidity0(amount0, price, maxPrice float64) float64 {
	return (amount0 * math.Sqrt(price) * math.Sqrt(maxPrice)) / (math.Sqrt(maxPrice) - math.Sqrt(price))
}

func GetLiquidity1(amount1, price, minPrice float64) float64 {
	return amount1 / (math.Sqrt(minPrice) - math.Sqrt(price))
}

// If you have an open position, calculate the liquidity0 open
func GetLiquidityWithAmount0(amount0, minPrice, maxPrice float64) float64 {
	a := math.Sqrt(minPrice) * math.Abs(math.Sqrt(maxPrice))
	b := math.Sqrt(maxPrice) - math.Sqrt(minPrice)
	return amount0 * (a / b)
}

// If you have an open position, calculate the liquidity1 open
func GetLiquidityWithAmount1(amount1, minPrice, maxPrice float64) float64 {
	quo := math.Sqrt(maxPrice) - math.Sqrt(minPrice)
	return amount1 / quo
}

// If you want to know the min price and have amount0, amount1, price and max Price
func GetMinPrice(amount0, amount1, price, maxPrice float64) float64 {
	sPrice := math.Sqrt(price)
	x := (amount1 / (math.Sqrt(maxPrice) * amount0)) + sPrice - (amount1 / (sPrice * amount0))
	return math.Pow(x, 2)
}
