"""
1.二分查找
时间复杂度：O(logn)
空间复杂度：O(1)
"""
class Solution:
    # 对有序序列的二分查找
    def binary_search(self, array, x):
        low, high = 0, len(array) - 1
        while low <= high:
            mid = (low + high) // 2
            if array[mid] == x:
                return mid
            elif array[mid] < x:
                low = mid + 1
            else:
                high = mid - 1
        return -1


if __name__ == "__main__":
    array = [1, 2, 4, 5, 9, 12]
    s = Solution()
    x = 5
    ret = s.binary_search(array, x)
    print(ret)
