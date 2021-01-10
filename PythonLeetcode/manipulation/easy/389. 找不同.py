"""

方法1： 相同字符进行异或运算则抵消为0，所以如果s比t少了一个字符，则直接将两个字符串的所有字符进行异或运算即可
时间复杂度：O(m+n)
空间复杂度：O(1)

case1:
abcd abcde
r:

"""
import sys
from typing import List


class Solution:
    @staticmethod
    def find_the_difference(s: str, t: str) -> str:
        ans = 0
        for i in (s + t):
            ans ^= ord(i)
        return chr(ans) if ans else ""


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        # 以后如果输入字符串要加上strip函数，否则末尾的"\n"会影响结果
        s_cur, t_cur = line.strip().split(" ")
        print(s.find_the_difference(s_cur, t_cur))
