"""
-11.字符串转整数.py
时间复杂度：O(1)
空间复杂度：O(1)
"""
import random


class Solution:
    def reservoir_sampling(self, array, k):
        if not array:
            return list()
        if k >= len(array):
            return array

        ret = array[:k]
        for i in range(k, len(array)):
            if random.randint(0, i) < k:
                ret[random.randint(0, k - 1)] = array[i]
        return ret


if __name__ == "__main__":
    array = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    k = 4
    s = Solution()
    ret = s.reservoir_sampling(array, k)
    print(ret)





