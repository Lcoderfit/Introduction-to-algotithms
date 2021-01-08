"""

方法1： 内置函数
时间复杂度：O(n)
空间复杂度：O(n)

方法2： 迭代贪心
时间复杂度：O(n)
空间复杂度：O(1)

case1:
27346209830709182346
r:
9

case1:
82734
r:
8

case2:
32
r:
3
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def min_partitions(n: str) -> int:
        return max(int(s) for s in n)

    @staticmethod
    def min_partitions2(n: str) -> int:
        ans = 0
        for s in n:
            if ans < int(s):
                ans = int(s)
        return ans


if __name__ == '__main__':
    sl = Solution()
    for line in sys.stdin:
        n_cur = line.strip()
        # print(sl.min_partitions(n_cur))
        print(sl.min_partitions2(n_cur))
