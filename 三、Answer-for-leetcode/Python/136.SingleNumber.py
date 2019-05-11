'''
1、位运算
'''

from typing import List


class Solution:
    def single_number(self, nums: List[int]) -> int:
        if len(nums) <= 0:
            raise Exception("nums is invalid!")

        res = 0
        for i in nums:
            res ^= i

        return res


if __name__ == '__main__':
    nums = list(map(int, input().split()))
    s = Solution()
    res = s.single_number(nums)
    print(res)
