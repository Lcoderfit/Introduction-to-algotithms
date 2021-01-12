"""

方法1： 
时间复杂度：O()
空间复杂度：O()

case1:
1 100 1 1 1 100 1 1 100 1
r:
6
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def min_costs_climbing_stairs(cost: List[int]) -> int:
        n = len(cost) + 1
        dp = [0] * n
        for i in range(2, n):
            dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
        return dp[n - 1]


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        cost_cur = [int(i) for i in line.strip().split(" ")]
        print(s.min_costs_climbing_stairs(cost_cur))
