"""
合并K个有序数组
算法时间复杂度：O(nlogK)
空间复杂度：O(KN)
"""


class Solution:
    def merge_sort_k(self, nums):
        if len(nums) <= 1:
            return nums

        mid = len(nums) // 2

        left = self.merge_sort_k(nums[:mid])
        right = self.merge_sort_k(nums[mid:])

        return self.sort_list(left[0], right[0])

    def sort_list(self, left, right):
        ret = []
        a, b = 0, 0

        while a < len(left) and b < len(right):
            if left[a] < right[b]:
                ret.append(left[a])
                a += 1
            else:
                ret.append(right[b])
                b += 1

        while a < len(left):
            ret.append(left[a])
            a += 1

        while b < len(right):
            ret.append(right[b])
            b += 1

        return [ret]


if __name__ == "__main__":
    nums = [[3, 4, 5], [1, 2, 3, 6], [4, 9, 11, 15, 17]]

    s = Solution()
    ret = s.merge_sort_k(nums)
    print(ret[0])
