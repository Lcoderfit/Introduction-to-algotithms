'''
1、位运算
'''
from typing import List


class Solution:
    def single_number(self, nums: List[int]) -> int:
        a, b = 0, 0
        for num in nums:
            b = ~a & (b ^ num)
            a = ~b & (a ^ num)

        return b


if __name__ == '__main__':
    while True:
        nums = list(map(int, input().split()))
        s = Solution()
        res = s.single_number(nums)
        print(res)

