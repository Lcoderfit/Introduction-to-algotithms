"""

方法1： 异或位运算
时间复杂度：O(1)
空间复杂度：O(1)

方法2： 取模和除法
时间复杂度：O(1)
空间复杂度：O(1)

case1:
10
r:
True

case2:
3
r:
False

case3:
17
r:
False

1 <= n <= 2**31 - 1

注：n&1 可以取到n二进制最低位的数，也可以判断奇偶性
"""
import sys
from typing import List


class Solution:
    # alternating: 交替、轮流
    @staticmethod
    def has_alternating_bits(n: int) -> bool:
        # 自己与右移后的自己异或
        t = (n ^ (n >> 1)) + 1
        return t & (t - 1) == 0

    @staticmethod
    # n&1 可以取到n二进制最低位的数，也可以判断奇偶性
    def has_alternating_bits2(n: int) -> bool:
        while n:
            if (n & 1) ^ ((n >> 1) & 1) == 0:
                return False
            n >>= 1
        return True


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        n_cur = int(line.strip())
        # res = s.has_alternating_bits(n_cur)
        res = s.has_alternating_bits2(n_cur)
        print(res)
