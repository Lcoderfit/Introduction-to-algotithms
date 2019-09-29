"""
14.减绳子
时间复杂度：O(1)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
class Solution:
    def cutRope1(self, number):
        # write code here
        if number <= 0:
            return 0
        if number == 1:
            return 1
        if number == 2:
            return 1
        if number == 3:
            return 2

        time3 = number // 3
        if number - time3 * 3 == 1:
            time3 -= 1
        time2 = (number - time3 * 3) // 2
        return (3 ** time3) * (2 ** time2)

    def cutRope2(self, number):
        if number <= 0:
            return 0
        if number == 1:
            return 1
        if number == 2:
            return 1
        if number == 3:
            return 2

        dp = [0] * (number + 1)
        dp[0] = 0
        dp[1] = 1
        dp[2] = 2
        dp[3] = 3
        for i in range(4, number + 1):
            for j in range(1, i//2 + 1):
                dp[i] = max(dp[i], dp[j] * dp[i - j])
        return dp[number]


if __name__ == "__main__":
    number = 10
    s = Solution()
    ret = s.cutRope2(number)
    print(ret)