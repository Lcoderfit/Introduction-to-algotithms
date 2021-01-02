"""
35. 搜索插入位置.py
时间复杂度：O(logn)
空间复杂度：O(1)
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def search_insert(nums: List[int], target: int) -> int:
        if not nums:
            return 0
        i, j = 0, len(nums) - 1
        while i <= j:
            mid = (i + j) // 2
            if target == nums[mid]:
                return mid
            elif nums[mid] > target:
                j = mid - 1
            else:
                i = mid + 1
        return i


if __name__ == '__main__':
    target_cur = int(input().strip())
    nums_cur = [int(i) for i in input().split(" ") if i.strip()]
    s = Solution()
    res = s.search_insert(nums_cur, target_cur)
    print(res)
