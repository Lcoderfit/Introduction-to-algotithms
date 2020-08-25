package dp

import "math"

func TwoSum(n int) []float64 {
	//dp[i][j]表示i个骰子掷出j点的次数
	dp := Create2DSlice(n+1, 6*n+1)
	//初始情况，只有一个骰子时，掷出1-6点的情况数都为1
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}

	//二维dp
	for i := 2; i <= n; i++ {
		//当有i个骰子时，能掷出的点数范围[i, 6*i]
		for j := i; j <= 6*i; j++ {
			//当前这个骰子可能掷出的点数为k==[1, 6]
			//dp[i][j]等于前一个骰子掷出j-k的点数的情况总和
			for k := 1; k <= 6; k++ {
				if k >= j {
					break
				}
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}
	//n个骰子所能掷出的总情况数
	sum := math.Pow(6, float64(n))
	ret := make([]float64, 0)
	for i := n; i <= 6*n; i++ {
		ret = append(ret, float64(dp[n][i])/sum)
	}
	return ret
}

//二维dp降为1维
func TwoSum2(n int) []float64 {
	dp := make([]int, 6*n+1)
	for i := 1; i <= 6; i++ {
		dp[i] = 1
	}

	for i := 2; i <= n; i++ {
		//如果j从小到大，先更新了dp[2],然后计算dp[3]时，
		//前面计算的dp[2]已经会对此时的dp[3]造成影响，也就是当前骰子数受当前骰子的影响，        //而原则上当前骰子点数只受前一个骰子点数影响，所以需要从大到小
		for j := 6*i; j >= i; j-- {
			//dp[j]初始值设为0
			dp[j] = 0
			for k := 1; k <= 6; k++ {
				//j-k为前i-1个骰子的点数和s, 而前i-1个骰子的最小点数为i-1
				if j - k < i - 1 {
					break
				}
				dp[j] += dp[j-k]
			}
		}
	}

	sum := math.Pow(6, float64(n))
	ret := make([]float64, 0)
	for i := n; i <= 6*n; i++ {
		ret = append(ret, float64(dp[i])/sum)
	}
	return ret
}