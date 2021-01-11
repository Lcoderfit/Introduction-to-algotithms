"""

方法1： 
时间复杂度：O()
空间复杂度：O()

case1:
1 2 3 4
r:
1 3 6 10

case2:
1 1 1 1 1
r:
1 2 3 4 5

case3:
3 1 2 10 1
r:
3 4 6 16 17
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def running_sum(nums: List[int]) -> List[int]:
        for i in range(1, len(nums)):
            nums[i] += nums[i - 1]
        return nums


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        nums_cur = [int(i) for i in line.strip().split(" ")]
        print(s.running_sum(nums_cur))
