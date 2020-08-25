"""
-5.最小的K个数
时间复杂度：O(nlogk)
空间复杂度：O(k)
"""
# -*- coding:utf-8 -*-
class Solution:
    def GetLeastNumbers_Solution(self, tinput, k):
        if not tinput:
            return list()
        if k <= 0 or k > len(tinput):
            return list()
        ret = tinput[:k]
        self.heapify(ret, 0, k - 1)
        for i in range(k, len(tinput)):
            if tinput[i] < ret[0]:
                ret[0] = tinput[i]
                self.heapify(ret, 0, k - 1)
        return sorted(ret)

    def heapify(self, array, start, end):
        root = start
        child = root * 2 + 1
        while child <= end:
            if child + 1 <= end and array[child] < array[child + 1]:
                child += 1
            if array[root] < array[child]:
                array[root], array[child] = array[child], array[root]
                root = child
                child = root * 2 + 1
            else:
                break


if __name__ == "__main__":
    tinput = [7, 2, 5, 1, 10, 4, 9]
    k = 3

    s = Solution()
    ret = s.GetLeastNumbers_Solution(tinput, k)
    print(ret)
