"""

TOPK问题的三种解法
一、冒泡法
时间复杂度：
空间复杂度：

二、堆排序
时间复杂度：O(kn)
空间复杂度：O(n)

三、快速排序
时间复杂度：
空间复杂度：
"""
from typing import List


# 方法一：冒泡排序
class Solution1:
    def bubbleSortK(self, array: List[int], k: int) -> List[int]:
        length = len(array)
        if k > length:
            return []

        count = 0
        for i in range(0, length - 1):
            count += 1
            if count == k:
                break
            for j in range(0, length - i -1):
                if array[j] > array[j + 1]:
                    array[j], array[j + 1] = array[j + 1], array[j]
        return array[0: k]


# 方法二：解决TOPK大的用小顶堆
class Solution2:
    # 小顶堆
    def heapify(self, arr, start, end):
        root = start
        child = root * 2 + 1
        while child <= end:
            if child + 1 <= end and arr[child] > arr[child + 1]:
                child += 1
            if arr[root] > arr[child]:
                arr[root], arr[child] = arr[child], arr[root]
                root = child
                child = root * 2 + 1
            else:
                break

    def topk(self, arr, k):
        ret = arr[:k]
        self.heapify(ret, 0, k - 1)
        for i in range(k, len(arr)):
            if arr[i] >= ret[0]:
                ret[0] = arr[i]
                self.heapify(ret, 0, k - 1)
        return ret[0]


# 方法三：快速排序
class Solution3:
    def largest_k(self, arr, low, high, k):
        if (k > 0 and k <= (high - low + 1)):
            index = self.partition(arr, low, high)
            if index - low + 1 == k:
                return arr[index]
            elif index - low + 1 > k:
                return self.largest_k(arr, low, index - 1, k)
            else:
                return self.largest_k(arr, index + 1, high, k - (index - low + 1))

    def partition(self, arr, low, high):
        pivot = arr[high]
        i = low
        for j in range(low, high):
            if arr[j] > pivot:
                arr[i], arr[j] = arr[j], arr[i]
                i += 1
        arr[i], arr[high] = arr[high], arr[i]
        return i


if __name__ == "__main__":
    array = [5, 3, 1, 2, 4, 9, 6, 8, 7]
    k = 6

    # 方法一：冒泡排序
    # s = Solution1()
    # ret = s.bubbleSortK(array, k)
    # print(ret)

    # 方法二：TOPK
    # s = Solution2()
    # ret = s.topk(array, 6)
    # print(ret)

    # 方法三：快速排序
    s = Solution3()
    ret = s.largest_k(array, 0, len(array) - 1, k)
    print(ret)