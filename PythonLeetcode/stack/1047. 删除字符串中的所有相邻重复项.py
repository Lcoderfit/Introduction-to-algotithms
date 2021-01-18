"""

方法1： 
时间复杂度：O(n)
空间复杂度：O(n)

case1:
abbaca
r:
ca
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def remove_duplicates(s: str) -> str:
        stack = []
        for ch in s:
            if stack and (stack[-1] == ch):
                stack.pop()
            else:
                stack.append(ch)
        return "".join(stack)


if __name__ == '__main__':
    sl = Solution()
    for line in sys.stdin:
        s_cur = line.strip()
        print(sl.remove_duplicates(s_cur))
