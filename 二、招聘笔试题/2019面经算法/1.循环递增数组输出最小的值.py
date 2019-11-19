"""
1.循环递增数组输出最小的值
时间复杂度：O(logn)
空间复杂度：O(1)
"""
class Solution:
    def minInCrcleArray(self, array):
        if not array:
            return None
        low = 0
        high = len(array) - 1
        ret = 0
        while low < high:
            mid = (low + high) // 2
            if array[mid] >= array[low] and array[mid] >= array[high]:
                low = mid + 1
            elif mid > 0 and array[mid] > array[mid - 1]:
                high = mid - 1
            else:
                ret = array[mid]
                break
        if low == high:
            ret = array[low]
        return ret


if __name__ == "__main__":
    array = [int(i) for i in input().split()]
    s = Solution()
    ret = s.minInCrcleArray(array)
    print(ret)
