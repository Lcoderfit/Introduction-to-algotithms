"""

方法1： 
时间复杂度：O()
空间复杂度：O()

case1:

r:

"""
import sys
from typing import List


class Solution:
    @staticmethod
    def add(a: int, b: int) -> int:
        a = (a & 0xFFFFFFFF)
        b = (b & 0xFFFFFFFF)

        while b:
            # 计算本位
            t = a
            a = t ^ b
            # 计算进位
            b = ((t & b) << 1) & 0xFFFFFFFF
            # 此时求到的a中，低位的32位为对应32位整型的补码表示，
            # 但是Python中有很多位，如果a >= 0x80000000; 则其对应的32位整型应该是个负数，
            # 而a此时高于32位的位应该均为0，需要做的事是保持低位32位不变，将高于32位的位全变为1
            # 则可以先将低位32位取反一次，然后所有位整体取反一次即可
        return a if a < 0x80000000 else ~(a ^ 0xFFFFFFFF)


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        a_cur, b_cur = [int(i) for i in line.strip().split(" ")]
        print(s.add(a_cur, b_cur))
