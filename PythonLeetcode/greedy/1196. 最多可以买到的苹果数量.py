"""

方法1： 模拟
时间复杂度：O(n)
空间复杂度：O(1)

case1:
100 200 150 1000
r:
4

case2:
900 950 800 1000 700 800
r:
5
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def max_number_of_apples(arr: List[int]) -> int:
        arr.sort()
        num = 0
        ans = 0
        for i in range(len(arr)):
            if num + arr[i] < 5000:
                num += arr[i]
                ans += 1
            else:
                break
        return ans


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        arr_cur = [int(i) for i in line.strip().split(" ")]
        print(s.max_number_of_apples(arr_cur))
