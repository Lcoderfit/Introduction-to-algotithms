from typing import List


class Solution:
    def majority_element(self, nums: List[int]) -> int:
        if len(nums) <= 0:
            raise Exception("nums is invalid!")

        count, res = 0, 0
        for num in nums:
            if count == 0:
                res = num
            if num == res:
                count += 1
            else:
                count -= 1
        return res


if __name__ == '__main__':
    while True:
        nums = list(map(int, input().split()))
        s = Solution()
        res = s.majority_element(nums)
        print(res)