package easy

import (
	"math"
)

//方法一
func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxVal := 0
	sum := 0
	for i := 1; i < len(prices); i++ {
		sum += prices[i] - prices[i-1]
		if sum < 0 {
			sum = 0
		} else {
			maxVal = Max(maxVal, sum)
		}
	}
	return maxVal
}

//方法二
func MaxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice, maxPrice := prices[0], prices[0]
	ret := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
			if ret < maxPrice-minPrice {
				ret = maxPrice - minPrice
			}
		}
		if prices[i] < minPrice {
			minPrice, maxPrice = prices[i], prices[i]
		}
	}
	return ret
}

//方法三
func MaxProfit3(prices []int) int {
	minPrice := int(math.MaxInt32)
	maxPrice := 0
	// 遍历数组，取prices[i]减去prices[0...i-1]中的最小值，然后取差值的最大值即可
	for _, price := range prices {
		maxPrice = Max(maxPrice, price-minPrice)
		minPrice = Min(minPrice, price)
	}
	return maxPrice
}
