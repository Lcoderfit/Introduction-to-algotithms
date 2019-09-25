# -*- coding:utf-8 -*-
"""
10-1.斐波那契堆
时间复杂度：O(n)
空间复杂度：O(n)
"""
class Solution:
    def Fibonacci(self, n):
        # write code here
        if n <= 0:
            return 0
        if n == 1:
            return 1
        dp = [0] * (n + 1)
        dp[1] = 1
        for i in range(2, n + 1):
            dp[i] = dp[i - 1] + dp[i - 2]
        return dp[n]

    def Fibonacci1(self, n):
        if n <= 0:
            return 0
        if n == 1:
            return 1
        return self.Fibonacci1(n - 1) + self.Fibonacci1(n - 2)


if __name__ == "__main__":
    s = Solution()
    ret = s.Fibonacci(5)
    print(ret)

    ret1 = s.Fibonacci1(6)
    print(ret1)