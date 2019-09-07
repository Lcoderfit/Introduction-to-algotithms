"""
选择排序：不稳定排序，
然后排序之后被交换到堆底，a之后也交换到堆底，则b在a的前面
算法时间复杂度
1、最优：O(n^2)
2、平均：O(n^2)
3、最差：O(n^2)
空间复杂度：O(1)
"""
from typing import List
import random


class Solution:
    def select_sort(self, collection):
        length = len(collection)
        for i in range(length - 1):
            least = i
            for k in range(i + 1, length):
                if collection[k] < collection[least]:
                    least = k
            collection[least], collection[i] = collection[i], collection[least]
        return collection

    def random_array(self, n):
        ret = []
        multi = set()
        for i in range(n):
            k = random.randint(1, n)
            while True:
                if k in multi:
                    k = random.randint(1, n)
                else:
                    multi.add(k)
                    ret.append(k)
                    break
        return ret


if __name__ == "__main__":
    s = Solution()
    arr = s.random_array(10)
    print(arr)
    ret = s.select_sort(arr)
    print(arr)