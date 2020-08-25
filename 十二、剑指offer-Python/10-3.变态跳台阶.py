"""
10-3.变态跳台阶
时间复杂度：O(n)
空间复杂度：O(n)
"""
# -*- coding:utf-8 -*-
class Solution:
    def jumpFloorII(self, number):
        # write code here
        if number <= 0:
            return 0
        elif number <= 2:
            return number
        dp = [0] * (number + 1)
        dp[1] = 1
        dp[2] = 2
        for i in range(3, number + 1):
            dp[i] = 2 * dp[i - 1]
        return dp[number]


if __name__ == "__main__":
    number = 5
    s = Solution()
    ret = s.jumpFloorII(number)
    print(ret)