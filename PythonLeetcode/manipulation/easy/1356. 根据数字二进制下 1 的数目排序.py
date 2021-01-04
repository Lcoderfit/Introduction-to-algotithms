"""

方法1： 内置函数
时间复杂度：O(nlogn)
空间复杂度：O(n)

方法2： TODO: 动态规划
时间复杂度：O(nlogn)
空间复杂度：O(n)

case1:
0 1 2 3 4 5 6 7 8
r1:
0 1 2 4 8 3 5 6 7

case2:
1024 512 256 128 64 32 16 8 4 2 1
r:
1 2 4 8 16 32 64 128 256 512 1024

case3:
10000 10000
r:
10000 10000

case4:
2 3 5 7 11 13 17 19
r:
2 3 5 17 7 11 13 19

case5:
10 100 1000 10000
r:
10 100 10000 1000
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def sort_by_bits(arr: List[int]) -> List[int]:
        return sorted(arr, key=lambda x: (bin(x).count("1"), x))


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        arr_cur = [int(i) for i in line.strip().split(" ")]
        print(s.sort_by_bits(arr_cur))
