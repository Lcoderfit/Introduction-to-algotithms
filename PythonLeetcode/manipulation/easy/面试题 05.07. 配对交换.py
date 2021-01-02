"""

方法1： 
时间复杂度：O()
空间复杂度：O()

case1:
5
r:
10

case2:
10
r:
5

注：取n的奇数位可以用n & 0x55555555, 取n的偶数位可以用0xaaaaaaaa
num的范围在[0, 2^30 - 1]之间，不会发生整数溢出
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def exchange_bits(num: int) -> int:
        return ((num & 0x55555555) << 1) | ((num & 0xaaaaaaaa) >> 1)


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        num_cur = int(line.strip())
        print(s.exchange_bits(num_cur))
