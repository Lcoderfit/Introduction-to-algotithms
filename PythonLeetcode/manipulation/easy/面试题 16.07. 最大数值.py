"""

方法1： 
时间复杂度：O()
空间复杂度：O()

case1:
1 2
r:
2

case1:
-1 1
r:
1
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def maximum(a: int, b: int) -> int:
        return (a + b + abs(a - b)) // 2

    @staticmethod
    def maximum2(a: int, b: int) -> int:
        c = a - b
        # 32位与32位进行加减，可能会超出范围，所以用64位长整型
        # 如果为负数，移位后为-1，如果为正数，移位后为0，+1即可转换为1和0
        sign = 1 + (c >> 63)
        # sign ^ 1，表示如果sign为1，则取0，sign为0，则取1
        return sign * a + (sign ^ 1) * b


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        a_cur, b_cur = [int(i) for i in line.strip().split(" ")]
        # print(s.maximum(a_cur, b_cur))
        print(s.maximum2(a_cur, b_cur))
