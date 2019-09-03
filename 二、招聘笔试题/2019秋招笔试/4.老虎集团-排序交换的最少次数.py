from typing import List


class Solution:
    def less_change(self, nums: List[int]) -> int:
        """使得数组有序的最少交换次数"""
        if not nums:
            return 0
        length = len(nums)
        if length == 1:
            return 0

        ans = 0
        for i in range(length):
            while nums[i] != (i + 1):
                temp = nums[i]
                nums[i] = nums[temp - 1]
                nums[temp - 1] = temp
                ans += 1

        return ans


if __name__ == "__main__":
    nums = [int(i) for i in input().split()]

    s = Solution()
    ans = s.less_change(nums)
    print(ans)