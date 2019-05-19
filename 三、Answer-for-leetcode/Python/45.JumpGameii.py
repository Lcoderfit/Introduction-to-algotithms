from typing import List


class Solution1:
    def jump(self, nums: List[int]) -> int:
        steps = last = curr = 0
        n = len(nums)
        for i in range(n):
            if i > last:
                steps += 1
                last = curr
            if curr < nums[i] + i:
                curr = nums[i] + i

        return steps


if __name__ == '__main__':
    while True:
        nums = list(map(int, input().split()))

        s = Solution1()
        res = s.jump(nums)
        print(res)