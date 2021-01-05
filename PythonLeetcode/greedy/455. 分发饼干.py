"""

方法1： 排序+贪心+双指针
时间复杂度：O(n) n = min(len(g), len(s))
空间复杂度：O(1)

case1:
1 2 3
1 1
r:
1

case2:
1 2
1 2 3
r:
2
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def find_content_children(g: List[int], s: List[int]) -> int:
        g.sort()
        s.sort()
        i, j = 0, 0
        ans = 0
        while (i < len(g)) and (j < len(s)):
            if g[i] <= s[j]:
                ans += 1
                i += 1
            j += 1
        return ans


if __name__ == '__main__':
    sl = Solution()
    g_cur = [int(i) for i in input().strip().split(" ")]
    s_cur = [int(i) for i in input().strip().split(" ")]
    print(sl.find_content_children(g_cur, s_cur))
