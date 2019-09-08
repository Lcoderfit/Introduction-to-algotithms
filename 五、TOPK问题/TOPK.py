"""
TOPK问题的三种解法
一、冒泡法
二、堆排序
三、快速排序
"""
from typing import List



class Solution:
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


if __name__ == "__main__":
    array = [5, 3, 1, 2, 4, 9, 6, 8, 7]
    k = 5

    s = Solution()
    ret = s.bubbleSortK(array, k)
    print(ret)