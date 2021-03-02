"""

方法1： 贪心
时间复杂度：O(n)
空间复杂度：O(1)

方法1： 动态规划
时间复杂度：O(n)
空间复杂度：O(n)

case1:
5
r:
13

case2:
3
r:
4
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def ways_to_step1(n: int) -> int:
        a, b, c = 1, 2, 4
        if n == 1:
            return a
        elif n == 2:
            return b
        elif n == 3:
            return c
        for i in range(4, n + 1):
            a, b, c = b, c, a + b + c
            if c > 1000000007:
                c = c % 1000000007
        return c

    @staticmethod
    def ways_to_step2(n: int) -> int:
        dp = [1, 2, 4] + [0] * (n - 3)
        for i in range(3, n):
            dp[i] = dp[i - 1] + dp[i - 2] + dp[i - 3]
            if dp[i] > 1000000007:
                dp[i] = dp[i] % 1000000007
        return dp[n - 1]


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        # res = s.ways_to_step1(int(line.strip()))
        res = s.ways_to_step2(int(line.strip()))
        print(res)
