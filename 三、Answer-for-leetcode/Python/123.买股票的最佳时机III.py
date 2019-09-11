# int max_k = 2;
# int[][][] dp = new int[n][max_k + 1][2];
# for (int i = 0; i < n; i++) {
#     for (int k = max_k; k >= 1; k--) {
#         if (i - 1 == -1) { /*处理 base case */ }
#         dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i]);
#         dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i]);
#     }
# }
# // 穷举了 n × max_k × 2 个状态，正确。
# return dp[n - 1][max_k][0];

class Solution:
    def max_profile(self, prices):
        max_k = 2
        dp = [[[0]*2]*max_k]*len(prices)
        for i in range(len(prices)):
            for k in range(max_k, 1, -1):
                if (i - 1 == -1):

