from typing import List


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        """滑动窗口的最大值，动态规划解法"""
        n = len(nums)
        if n * k == 0:
            return list()
        if k == 1:
            return nums

        left = [0] * n
        left[0] = nums[0]
        right = [0] * n
        right[0] = nums[n - 1]
        for i in range(1, n - k - 1):
            if i % k == 0:
                left[i] = nums[i]
            else:
                left[i] = max(left[i - 1], nums[i])

            j = n - i - 1
            if (j + 1) % k == 0:
                right[j] = nums[j]
            else:
                right[j] = max(right[j + 1], nums[j])

        ret = []
        for i in range(n - k + 1):
            ret.append(max(left[i + k -1], right[i]))

        return ret


class Solution1:
    def maxSlidingWindow(self, nums, k):
        if not nums:
            return list()
        if k == 1:
            return nums

        win, ret = [], []
        for k, v in enumerate(nums):
            if i >= k and i - k >= win[0]:
                win.pop(0)
            while win and win[-1] < v:
                win.pop()
            win.append(i)
            if i >= k - 1:
                ret.append(nums[win[0]])
        return ret
