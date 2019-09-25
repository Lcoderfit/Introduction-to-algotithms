"""
归并排序：稳定排序，当只有一个元素的时候退出递归，两个相等的元素拆开后合并，原来的顺序没变
算法时间复杂度
T[n] = T[n/2] + O(n)
1、最优：O(nlogn)
2、平均：O(nlogn)
3、最差：O(nlogn)
空间复杂度：O(n)
"""
import random


class Solution:
    def merge_sort(self, collection):
        length = len(collection)
        if length > 1:
            midpoint = length // 2
            left_half = self.merge_sort(collection[:midpoint])
            right_half = self.merge_sort(collection[midpoint:])
            i, j, k = 0, 0, 0
            left_length = len(left_half)
            right_length = len(right_half)

            while i < left_length and j < right_length:
                if left_half[i] < right_half[j]:
                    collection[k] = left_half[i]
                    i += 1
                else:
                    collection[k] = right_half[j]
                    j += 1

                k += 1

            while i < left_length:
                collection[k] = left_half[i]
                i += 1
                k += 1

            while j < right_length:
                collection[k] = right_half[j]
                j += 1
                k += 1

        return collection

    def random_array(self, n):
        array = list()
        multi = set()
        for i in range(n):
            k = random.randint(1, n)
            while True:
                if k in multi:
                    k = random.randint(1, n)
                else:
                    multi.add(k)
                    array.append(k)
                    break
        return array


if __name__ == "__main__":
    # array = [int(item) for item in input().split()]

    s = Solution()
    array = s.random_array(10)
    print(array)
    ret = s.merge_sort(array)

    print(ret)
