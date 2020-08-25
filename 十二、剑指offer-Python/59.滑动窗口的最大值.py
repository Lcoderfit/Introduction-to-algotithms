"""
59.滑动窗口的最大值
时间复杂度：O(n)
空间复杂度：O(n)
"""
# -*- coding:utf-8 -*-
class Solution:
    def maxInWindows(self, num, size):
        if size < 1 or size > len(num):
            return list()

        index = list()
        for i in range(size):
            while index and num[i] >= num[index[-1]]:
                index.pop()
            index.append(i)

        ret = list()
        for i in range(size, len(num)):
            ret.append(num[index[0]])

            while index and num[i] >= num[index[-1]]:
                index.pop()
            if index and index[0] <= (i - size):
                index.pop(0)

            index.append(i)
        ret.append(num[index[0]])

        return ret


if __name__ == "__main__":
    num = [2, 3, 4, 2, 6, 2, 5, 1]
    size = 3
    s = Solution()
    ret = s.maxInWindows(num, 3)
    print(ret)