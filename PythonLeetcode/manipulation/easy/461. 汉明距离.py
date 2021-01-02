"""

方法1： 布莱恩.克尼根算法， n & (n-1)可以快速消除n二进制表示最右边为1的一位
时间复杂度：O(1), 整型长度固定，最多移位32次
空间复杂度：O(1)

方法2： 内置位计算功能
时间复杂度：O(1)
空间复杂度：O(1)

方法3： 移位
时间复杂度：O(1)
空间复杂度：O(1)

case1:
1 5
r1:
1

case2:
3 10
r1:
2

注意：
0 ≤ x, y < 2^31.
否则如果x，y为一正一负，x与y异或之后，最高位的表示符号位的1会一直存在，陷入无限循环
"""
import sys
from typing import List


class Solution:
    # distance：间距，距离
    # xor： 异或,exclusive or
    @staticmethod
    def hamming_distance(x: int, y: int) -> int:
        xor = x ^ y
        distance = 0
        while xor:
            distance += 1
            xor = xor & (xor - 1)
        return distance

    @staticmethod
    def hamming_distance2(x: int, y: int) -> int:
        return bin(x ^ y).count("1")

    @staticmethod
    def hamming_distance3(x: int, y: int) -> int:
        xor = x ^ y
        distance = 0
        while xor:
            if xor & 1:
                distance += 1
            xor = xor >> 1
        return distance


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        x_cur, y_cur = [int(i) for i in line.split(" ")]
        # res = s.hamming_distance(x_cur, y_cur)
        # res = s.hamming_distance2(x_cur, y_cur)
        res = s.hamming_distance3(x_cur, y_cur)
        print(res)
