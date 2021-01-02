"""

方法1： 异或位运算
时间复杂度：O(n)
空间复杂度：O(1)

方法2： 内置函数
时间复杂度：O(n)
空间复杂度：O(1)

case1:
3 0 1
r:
2

case2:
9 6 4 2 3 5 7 0 1
r:
8
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def missing_number(nums: List[int]) -> int:
        # 利用相同的数字异或为0的原理，将列表下标与元素相异或
        ans = len(nums)
        for i in range(len(nums)):
            ans = ans ^ i ^ nums[i]
        return ans

    @staticmethod
    def missing_number2(nums: List[int]) -> int:
        return sum(range(len(nums) + 1)) - sum(nums)

    @staticmethod
    def missing_numbers3(nums: List[int]) -> int:
        ans = len(nums)
        for i in range(len(nums)):
            ans += i - nums[i]
        return ans


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        nums_cur = [int(i) for i in line.split(" ")]
        res = s.missing_number(nums_cur)
        print(res)
