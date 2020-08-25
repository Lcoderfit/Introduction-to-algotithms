package dp

import "math"

func MaxProfit(prices []int) int {
	minPrice, maxPrice := math.MaxInt32, 0
	for _, v := range prices {
		maxPrice = max(maxPrice, v-minPrice)
		minPrice = min(minPrice, v)
	}
	return maxPrice
}