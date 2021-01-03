"""

方法1： 双指针
时间复杂度：O(1)
空间复杂度：O(1)

方法2： 字符串分隔
时间复杂度：O(1)
空间复杂度：O(1)

case1:
1775
r:
8

case2:
-1
r:
32

case3:
-2
r:
32

case4:
7
r:
4

num为一个32位整数
统计类问题用双指针，或者双指针配合循环
"""
import sys
from typing import List


# TODO: 双指针妙用
class Solution:
    @staticmethod
    def reverse_bits(num: int) -> int:
        ans = 0
        count = 1
        pos = -1
        num = num & 0xFFFFFFFF
        for i in range(32):
            if num & 1:
                count += 1
            else:
                ans = max(ans, count)
                count = i - pos
                pos = i
            num = (num >> 1)
        ans = max(ans, count)
        return ans if ans <= 32 else 32

    @staticmethod
    def reverse_bits2(num: int) -> int:
        ans = 0
        num = num & 0xFFFFFFFF
        bits_list = bin(num)[2:].split("0")
        if len(bits_list) <= 1:
            if len(bits_list[0]) == 32:
                return 32
            return len(bits_list[0]) + 1
        for i in range(len(bits_list) - 1):
            ans = max(ans, len(bits_list[i]) + len(bits_list[i + 1]))
        return ans + 1


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        num_cur = int(line.strip())
        # print(s.reverse_bits(num_cur))
        print(s.reverse_bits2(num_cur))

    # # 如果只有一个连续的“0”，则会将0前后分为两部分，如果有多个连续的“0”
    # # 则除了分割的“0”外，多出来的“0”均将在生成的列表中以空字符形式存在.
    # print("10000111110111".split("0"))
