"""

方法1： 贪心法 （while版）
时间复杂度：O(n)
空间复杂度：O(1)

方法2： 贪心法 (for版)
时间复杂度：O(n)
空间复杂度：O(1)

case1:
RLRRLLRLRL
r:
4

case2:
RLLLLRRRLR
r:
3

case3:
LLLLRRRR
r:
1
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def balanced_string_split(s: str) -> int:
        ans, i = 0, 0
        balance = 0
        while i < len(s):
            if s[i] == 'R':
                balance += 1
            else:
                balance -= 1
            if balance == 0:
                ans += 1
            i += 1
        return ans

    @staticmethod
    def balanced_string_split2(s: str) -> int:
        ans = 0
        balance = 0
        for i in s:
            if i == "R":
                balance += 1
            else:
                balance -= 1
            if balance == 0:
                ans += 1
        return ans


if __name__ == '__main__':
    sl = Solution()
    for line in sys.stdin:
        s_cur = line.strip()
        print(sl.balanced_string_split(s_cur))
