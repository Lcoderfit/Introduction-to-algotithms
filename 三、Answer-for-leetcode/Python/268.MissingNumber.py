from typing import List


class Solution1:
    '''用set求差集'''
    def missing_number(self, nums: List[int]) -> int:
        return list(set(range(len(nums) + 1)) - set(nums))[0]


class Solution2:
    '''用sum求和的差'''
    def missing_number(self, nums: List[int]) -> int:
        return sum(range(len(nums)+1)) - sum(nums)


if __name__ == '__main__':
    while True:
        nums = list(map(int, input().split()))
        s = Solution1()
        res = s.missing_number(nums)
        print(res)