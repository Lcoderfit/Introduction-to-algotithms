"""

方法1： 
时间复杂度：O(n)
空间复杂度：O(n)

case1:

r:

"""
import sys
from typing import List


class Solution:
    @staticmethod
    def make_good(s: str) -> str:
        stack = []
        for ch in s:
            if stack and abs(ord(ch) - ord(stack[-1])) == 32:
                stack.pop()
            else:
                stack.append(ch)
        return "".join(stack)


if __name__ == '__main__':
    sl = Solution()
    for line in sys.stdin:
        s_cur = line.strip()
        print(sl.make_good(s_cur))
