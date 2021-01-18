"""

方法1： 
时间复杂度：O(n)
空间复杂度：O(n)

case1:
({[()()]}){}{}][
r:
False

case2:
()[]{}
r:
True

case3:
([)]
r:
False

case4:
{[]}
False:
True

栈类型的题目有明显的向左退位性质
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def is_valid(s: str) -> bool:
        if len(s) & 1:
            return False

        stack = []
        pairs = {
            ")": "(",
            "]": "[",
            "}": "{"
        }

        for ch in s:
            if ch in pairs:
                if (not stack) or (stack[-1] != pairs[ch]):
                    return False
                stack.pop()
            else:
                stack.append(ch)
        return not stack


if __name__ == '__main__':
    sl = Solution()
    for line in sys.stdin:
        s_cur = line.strip()
        print(sl.is_valid(s_cur))
