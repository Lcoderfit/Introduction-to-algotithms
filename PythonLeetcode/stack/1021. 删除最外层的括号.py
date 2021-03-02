"""

方法1： 
时间复杂度：O()
空间复杂度：O()

case1:
(()())()(())
r:
()()()()

case2:
()()
r:
""
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def remove_outer_parentheses(s: str) -> str:
        if not s:
            return ""
        stack = [s[0]]
        pos = 0
        ans = ""
        for i in range(1, len(s)):
            if not stack:
                ans += s[pos + 1:i - 1]
                pos = i
                stack.append(s[i])
            else:
                if stack[-1] == "(" and s[i] == ")":
                    stack.pop()
                else:
                    stack.append(s[i])
        return ans + s[pos + 1:-1]


if __name__ == '__main__':
    sl = Solution()
    s_cur = input().strip()
    print(sl.remove_outer_parentheses(s_cur))
