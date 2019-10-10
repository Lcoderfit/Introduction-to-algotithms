"""
45.把数组排成最小的数.py
时间复杂度：O(n)
空间复杂度：O(n)
"""
# -*- coding:utf-8 -*-
class LargeNumKey(str):
    def __lt__(self, other):
        return self + other < other + self  # 要使”32“< "3"返回True，则"323" < "332"应返回True


class Solution:
    def PrintMinNumber(self, numbers):
        # write code here
        if not numbers:
            return ""
        ret = sorted(map(str, numbers), key = LargeNumKey)
        return int("".join(ret))


from functools import cmp_to_key
import random


class Solution2:
    def PrintMinNumber(self, numbers):
        # write code here
        if not numbers:
            return ""
        temp = list(map(str, numbers))
        print(temp)
        ret = sorted(temp, key = cmp_to_key(lambda x, y: int(x + y) - int(y + x)))
        return "".join(ret if ret[0] != "0" else "0")


if __name__ == "__main__":
    numbers = [3, 5, 1, 4, 2]

    s = Solution2()
    ret = s.PrintMinNumber(numbers)
    print(ret)