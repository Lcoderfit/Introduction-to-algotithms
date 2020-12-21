"""

方法1： hash表
时间复杂度：O(n)
空间复杂度：O(n)

方法2： 直接遍历
时间复杂度：O(n)
空间复杂度：O(1)

方法3： 双边界二分查找
时间复杂度：O(logn)
空间复杂度：O(1)

方法4： 调用函数
时间复杂度：O(?)
空间复杂度：O(?)

方法5： 单次二分查找
时间复杂度：O(logn)
空间复杂度：O(1)

case1:
5
2 4 5 5 5 5 5 6 6

case2:
101
10 100 101 101
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def is_majority_element(nums: List[int], target: int) -> bool:
        map_index = dict()
        for i in range(len(nums)):
            if nums[i] in map_index:
                map_index[nums[i]] += 1
            else:
                map_index[nums[i]] = 1
        border = len(nums) // 2
        if (target in map_index) and (map_index[target] > border):
            return True
        return False

    @staticmethod
    def is_majority_element2(nums: List[int], target: int) -> bool:
        count = 0
        for i in range(len(nums)):
            if nums[i] == target:
                count += 1
        return count * 2 > len(nums)


    @staticmethod
    def is_majority_element3(nums: List[int], target: int) -> bool:
        """先找到左边界，再找到右边界"""
        i, j = 0, len(nums) - 1
        while i < j:
            mid = (i + j) // 2
            if nums[mid] >= target:
                j = mid
            else:
                i = mid + 1
        if nums[i] != target:
            return False
        left = i

        # 如果是求右边界，则mid偏向右，nums[mid]与nums元素值相等的情况时i指向mid，不偏移，因为mid会偏向右
        # 右边界，mid偏右，左指指向mid
        # 如果循环条件为i <= j,则可以改成i向mid右偏，j向mid左偏
        i, j = 0, len(nums) - 1
        while i < j:
            mid = (i + j + 1) // 2
            if nums[mid] <= target:
                i = mid
            else:
                j = mid - 1
        right = i
        return (right - left + 1) > len(nums) // 2

    @staticmethod
    def is_majority_element4(nums: List[int], target: int) -> bool:
        return nums.count(target)*2 > len(nums)

    def is_majority_element5(self, nums: List[int], target: int) -> bool:
        left = self.binary_search_left(nums, target)
        return (left != -1) and (left + len(nums) // 2 < len(nums)) and (nums[left + len(nums) // 2] == target)

    @staticmethod
    def binary_search_left(nums: List[int], target: int) -> int:
        # i <= j与 i = mid + 1; j = mid - 1配合
        # i < j 与 i = mid + 1; j = mid配合
        i, j = 0, len(nums) - 1
        while i <= j:
            mid = (i + j) // 2
            if nums[mid] >= target:
                j = mid - 1
            else:
                i = mid + 1
        if (i < len(nums)) and (nums[i] == target):
            return i
        return -1


if __name__ == '__main__':
    target_cur = int(input().strip().split(" ")[0])
    nums_cur = [int(i) for i in input().strip().split(" ")]
    s = Solution()
    # res = s.is_majority_element(nums_cur, target_cur)
    res = s.is_majority_element2(nums_cur, target_cur)
    print(res)
