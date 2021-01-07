"""

方法1： 
时间复杂度：O(nlogn)
空间复杂度：O(1)

case1:
4 3 10 9 8
r:
10 9

case2:
4 4 7 6 7
r:
7 7 6

case3:
6
r:
6
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def min_sub_sequence(nums: List[int]) -> List[int]:
        nums.sort(reverse=True)
        sum_value = sum(nums)
        ans = []
        ans_sum = 0
        for v in nums:
            ans_sum += v
            ans.append(v)
            if ans_sum > sum_value - ans_sum:
                break
        return ans


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        nums_cur = [int(i) for i in line.strip().split(" ")]
        print(s.min_sub_sequence(nums_cur))
