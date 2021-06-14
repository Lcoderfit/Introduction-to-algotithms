package medium

import "math"

func ChampagneTower(poured int, queryRow int, queryGlass int) float64 {
	dp := make([][]float64, 102)
	for i := range dp {
		dp[i] = make([]float64, 102)
	}
	// 第一杯初始化为倒入的所有香槟数
	dp[0][0] = float64(poured)
	// 可以假设第一杯先装下所有的香槟（设为n），如果比1大，则会溢出，溢出量为：n-1， 溢出的香槟会平分给与之相邻的两个香槟，
	// 于是dp[i+1][j+1]和dp[i+1][j]分别加上 n-1/2.0
	for i := 0; i < queryRow; i++ {
		for j := 0; j <= i; j++ {
			q := (dp[i][j] - 1.0) / 2.0
			if q > 0 {
				dp[i+1][j] += q
				dp[i+1][j+1] += q
			}
		}
	}
	return math.Min(1.0, dp[queryRow][queryGlass])
}
