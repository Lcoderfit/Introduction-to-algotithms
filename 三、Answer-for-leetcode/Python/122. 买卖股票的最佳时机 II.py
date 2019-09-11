"""
122.买股票的最佳时期II：O(n)
7 1 5 3 6 4
考虑[1, 2, 3]：最大利润3 - 1，与(2 - 1) + (3 - 2)相等，所以最大利润 = 每次升序的差值之和
"""
from typing import List


class Solution:
    def maxProfile(self, prices: List[int]) -> int:
        if len(prices) <= 1:
            return 0
        max = 0
        for i in range(1, len(prices)):
            div = prices[i] - prices[i - 1]
            max += div if div > 0 else 0    #必须是div>0,因为div有可能为负数

        return max


if __name__ == "__main__":
    prices = [int(i) for i in input().split()]

    s = Solution()
    ret = s.maxProfile(prices)
    print(ret)
