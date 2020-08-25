"""
49.丑数.py
时间复杂度：O(n)
空间复杂度：O(n)
"""
# -*- coding:utf-8 -*-
class Solution:
    def GetUglyNumber_Solution(self, index):
        # write code here
        if index <= 0:
            return 0
        if index == 1:
            return 1
        ret = [0] * (index + 1)
        ret[1] = 1
        t2, t3, t5 = 1, 1, 1
        for i in range(2, index + 1):
            while ret[t2] * 2 <= ret[i - 1]:
                t2 += 1
            while ret[t3] * 3 <= ret[i - 1]:
                t3 += 1
            while ret[t5] * 5 <= ret[i - 1]:
                t5 += 1
            k2 = ret[t2] * 2
            k3 = ret[t3] * 3
            k5 = ret[t5] * 5
            ret[i] = min(k2, k3, k5)
        return ret[index]


if __name__ == "__main__":
    index = 10
    s = Solution()
    ret = s.GetUglyNumber_Solution(index)
    print(ret)
    a = "string"
    b = 'k'
    print(ord)

