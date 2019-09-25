"""
10-4.覆盖矩形
时间复杂度：O(n)
空间复杂度：O(n)
"""
# -*- coding:utf-8 -*-
class Solution:
    def rectCover(self, number):
        # write code here
        if number <= 0:
            return 0
        if number == 1:
            return 1
        if number == 2:
            return 2
        dp = [0] * (number + 1)
        dp[1] = 1
        dp[2] = 2
        for i in range(3, number + 1):
            dp[i] = dp[i - 1] + dp[i - 2]
        return dp[number]


if __name__ == "__main__":
    number = 6
    s = Solution()
    ret = s.rectCover(number)
    print(ret)