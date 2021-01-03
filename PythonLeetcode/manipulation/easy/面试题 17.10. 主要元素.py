"""

方法1： 摩尔投票法
时间复杂度：O(n)
空间复杂度：O(1)

方法2： 位运算
时间复杂度：O(n)
空间复杂度：O(1)

case1:
1 2 5 9 5 9 5 5 5
r:
5

case2:
1 2 3
r:
-1

case3:
1 2 2 3
r1
-1
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def majority_element(nums: List[int]) -> int:
        major = 0
        ans = 0
        for i in range(len(nums)):
            if major == 0:
                ans = nums[i]
            if ans == nums[i]:
                major += 1
            else:
                major -= 1

        count = 0
        for i in range(len(nums)):
            if ans == nums[i]:
                count += 1
        return ans if count > len(nums) // 2 else -1


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        nums_cur = [int(i) for i in line.strip().split(" ")]
        print(s.majority_element(nums_cur))
