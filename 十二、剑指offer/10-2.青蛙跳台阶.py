# -*- coding:utf-8 -*-
"""
10-2.青蛙跳台阶
时间复杂度：O(n)
空间复杂度：O(n)
"""
class Solution:
    def jumpFloor(self, number):
        # write code here
        dp = [0] * (number + 1)
        dp[0] = 1
        dp[1] = 1
        if number <= 0:
            return 0
        for i in range(2, number + 1):
            dp[i] = dp[i - 1] + dp[i - 2]

        return dp[number]