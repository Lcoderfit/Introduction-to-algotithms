from typing import List


class Solution(object):
    def single_number(self, nums: List[int]) -> tuple:
        diff = 0
        for num in nums:
            diff ^= num
        filter = diff & ~(diff - 1)
        a, b = 0, 0
        for num in nums:
            if num & filter:
                a ^= num
            else:
                b ^= num
        return a, b


if __name__ == '__main__':
    while True:
        nums = list(map(int, input().split()))
        s = Solution()
        res = s.single_number(nums)
        print(res)