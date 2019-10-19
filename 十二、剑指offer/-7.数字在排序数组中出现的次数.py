"""
-7.数字在排序数组中出现的次数.py
时间复杂度：O(logn)
空间复杂度：O(1)
"""
class Solution:
    def GetNumberOfK(self, data, k):
        if not data:
            return 0
        left, right = 0, len(data) - 1
        leftk = self.GetLeftK(data, k, left, right)
        rightk = self.GetRightK(data, k, left, right)

        return rightk - leftk + 1

    def GetLeftK(self, data, k, left, right):
        while left <= right:
            mid = (left + right) // 2
            if data[mid] < k:
                left += 1
            else:
                right -= 1
        return left

    def GetRightK(self, data, k, left, right):
        while left <= right:
            mid = (left + right) // 2
            if data[mid] <= k:
                left += 1
            else:
                right -= 1
        return right


if __name__ == "__main__":
    data = [1, 3, 4, 4, 4, 5, 6, 9]
    k = 4
    s = Solution()
    ret = s.GetNumberOfK(data, k)
    print(ret)
