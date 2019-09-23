"""
53.最大子序和
时间复杂度：O(n)
空间复杂度：O(n)
"""
from typing import List


class Solution:
    def max_sub_array(self, nums: List[int]) -> int:
        ret, sum = nums[0], nums[0]
        for i in range(1, len(nums)):
            sum = max(sum + nums[i], nums[i])
            if ret < sum:
                ret = sum

        return ret


if __name__ == "__main__":
    # nums = [int(i) for i in input().split(",")]
    nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
    s = Solution()
    ret = s.max_sub_array(nums)
    print(ret)



