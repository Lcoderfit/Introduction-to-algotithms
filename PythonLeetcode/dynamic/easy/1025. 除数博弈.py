"""

方法1： 动态规划
时间复杂度：O(nlogn)
空间复杂度：O(n)

方法1： 数学
时间复杂度：O(1)
空间复杂度：O(1)

case1:
10
r:
true

"""
import math
import sys
from typing import List


class Solution:
    @staticmethod
    def divisor_game(n: int) -> bool:
        dp = [False] * (n + 1)
        for i in range(1, n + 1):
            # math.ceil向上取整，math.floor向下取整
            border = math.ceil(math.sqrt(i))
            for j in range(1, border):
                # if i % j != 0:
                #     continue
                # if not dp[i - j]:
                #     dp[i] = True
                #     break
                if (i % j != 0) and (not dp[i - j]):
                    dp[i] = True
                    break
        return dp[n]

    @staticmethod
    def divisor_game1(n: int) -> bool:
        # 奇数必败，偶数必胜
        return n & 1 != 1


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        # res = s.divisor_game(int(line.strip()))
        res = s.divisor_game1(int(line.strip()))
        print(res)
