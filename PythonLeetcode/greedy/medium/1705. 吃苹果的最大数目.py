"""

方法1： 
时间复杂度：O()
空间复杂度：O()

case1:
1 2 3 5 2
3 2 1 4 2
r:
7

case2:
3 0 0 0 0 2
3 0 0 0 0 2
r:
5
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def eaten_apples(apples: List[int], days: List[int]) -> int:
        ans, day = 0, 1
        for i in range(1, len(apples) + 1):
            if day < i:
                day = i
            if day >= i + days[i - 1]:
                continue
            apple_count = min(apples[i - 1], i + days[i - 1] - day)
            ans += apple_count
            day += apple_count
        return ans


if __name__ == '__main__':
    s = Solution()
    apples_cur = [int(i) for i in input().strip().split(" ")]
    days_cur = [int(i) for i in input().strip().split(" ")]
    print(s.eaten_apples(apples_cur, days_cur))
