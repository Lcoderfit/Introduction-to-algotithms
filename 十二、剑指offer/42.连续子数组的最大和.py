"""
-4.连续子数组的最大和
时间复杂度：O(n)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
class Solution:
    def FindGreatestSumOfSubArray(self, array):
        # write code here
        sum, ret = array[0], array[0]
        for i in range(1, len(array)):
            # dp[i] = max(dp[i-1] + array[i], array[i])
            sum  = max(sum + array[i], array[i])
            if ret < sum:
                ret = sum
        return ret


if __name__ == "__main__":
    array = [6,-3,-2,7,-15,1,2,2]
    s = Solution()
    ret = s.FindGreatestSumOfSubArray(array)
    print(ret)