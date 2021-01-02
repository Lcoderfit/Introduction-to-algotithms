"""

方法1： 内置位计算功能
时间复杂度：O(1)
空间复杂度：O(1)

方法2： 递归
时间复杂度：O(1)
空间复杂度：O(1)
正整数判断奇偶直接用位运算：n & 1

case1:
5
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def number_of_steps(num: int) -> int:
        num = bin(num)
        return len(num) + num.count("1") - 3

    # @staticmethod
    # def number_of_steps2(num: int) -> int:

    def number_of_steps2(self, num: int) -> int:
        if num == 0:
            return 0
        # 如果是奇数，则-1操作可以转换为将二进制表示中最后一位从1转换为0，在Python中，即与0xFFFF...FFFFFFe(28字节)进行&操作
        # 而0xFFFF...FFFFFFe对应Python整型-2的补码
        return 1 + self.number_of_steps2(num - 1 if num & 1 else num >> 1)


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        num_cur = int(line.strip())
        print(s.number_of_steps(num_cur))
        # print(s.number_of_steps2(num_cur))

        # # x不论为1还是0，与1相与都会保持其原有的样子, 所以num_cur与负数相与，其正负性不会受到影响，因为-2高位均为1
        # print(num_cur & -2)
