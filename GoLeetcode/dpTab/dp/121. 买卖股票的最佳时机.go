package dp

import "math"

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
			maxVal = max(maxVal, sum)
		}
	}
	return maxVal
}

//方法二
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice, maxPrice := prices[0], prices[0]
	ret := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
			if ret < maxPrice - minPrice {
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
func MaxProfit2(prices []int) int {
	minPrice := int(math.MaxInt32)
	maxPrice := 0
	for _, price := range prices {
		maxPrice = max(maxPrice, price - minPrice)
		minPrice = min(minPrice, price)
	}
	return maxPrice
}