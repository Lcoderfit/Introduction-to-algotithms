"""
1.循环递增数组输出最小的值
时间复杂度：O(logn)
空间复杂度：O(1)
"""
class Solution:
    def minInCircleArray(self, array):
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

    def minInCircleArrayRecursion(self, array, begin, end):
        if not array:
            return None
        # 如果array[begin] == array[end]: array中只有一个元素
        # 如果array[begin] < array[end]: 说明array中元素是按照递增排列的
        if array[begin] <= array[end]:
            return array[begin]
        # 数组只有两个元素时
        if end - begin == 1:
            return array[end]

        mid = (begin + end) // 2
        # 刚好是最小元素
        if array[mid] < array[mid - 1]:
            return array[mid]

        if array[mid] > array[begin]:
            return self.minInCircleArrayRecursion(array, mid + 1, end)
        else:
            return self.minInCircleArrayRecursion(array, begin, mid - 1)


if __name__ == "__main__":
    while True:
        array = [int(i) for i in input().split()]
        s = Solution()
        ret = s.minInCircleArray(array)
        print("minInCircleArray:", ret)

        retRecursion = s.minInCircleArrayRecursion(array, 0, len(array) - 1)
        print("minInCircleArrayRecursion:", retRecursion)
