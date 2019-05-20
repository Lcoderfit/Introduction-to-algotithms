from typing import List


class Solution1:
    def can_jump(self, nums: List[int]) -> bool:
        test = 0
        for i in range(len(nums)-1):
            test = max(test, i+nums[i])
            if test <= i:
                return False
        return True


class Solution2:
    def can_jump(self, nums: List[int]) -> bool:
        n = len(nums)
        l = n - 1
        for i in range(n-1, -1, -1):
            if i + nums[i] >= l:
                l = i
        return l <= 0


if __name__ == '__main__':
    while True:
        nums = list(map(int, input().split()))

        s = Solution1()
        res = s.can_jump(nums)
        print(res)