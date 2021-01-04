"""

方法1： 字符串操作（自己写的shit）
时间复杂度：O(1)
空间复杂度：O(1)

方法2： 位运算
时间复杂度：O(1)
空间复杂度：O(1)

case1:
-100
r1:
ffffff9c

case2:
0
r1:
0

case3:
154
r1:
9a
"""
import sys
from typing import List


class Solution:
    def to_hex(self, num: int) -> str:
        ans = ""
        bin_str = bin(num & 0xFFFFFFFF)[2:]
        for i in range(len(bin_str), -1, -4):
            left = i - 4 if i - 4 >= 0 else 0
            ans = self.bin_to_hex(bin_str[left:i]) + ans
        return ans.lstrip("0") if ans != "0" else "0"

    @staticmethod
    def bin_to_hex(bin_str):
        bin_str = bin_str[::-1]
        res = 0
        base = 2
        for i in range(0, len(bin_str)):
            if bin_str[i] == "1":
                res += base ** i
        if res > 9:
            return chr(res + 87)
        return str(res)

    @staticmethod
    def to_hex2(num):
        ans = ""
        # 利用字符串与其索引之间的关系
        hex_str = "0123456789abcdef"
        num = num & 0xFFFFFFFF
        if not num:
            return "0"
        while num:
            ans = hex_str[num % 16] + ans
            num = (num >> 4) & 0xFFFFFFFF
        return ans


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        num_cur = int(line.strip())
        # print(s.to_hex(num_cur))
        print(s.to_hex2(num_cur))
        # print(s.bin_to_hex(num_cur))
