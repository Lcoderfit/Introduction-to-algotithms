"""
16.数值的整数次方
时间复杂度：如果指数为n，则算法时间复杂度为O(n), 有一种高效解法为O(logn)
空间复杂度：O(1)
"""
import random


# -*- coding:utf-8 -*-
class Solution:
    def Power(self, base, exponent):
        # write code here
        if base == 0.0 and exponent <= 0:
            return 0.0

        ret = 1.0
        t_exponent = abs(int(exponent))
        for i in range(t_exponent):
            ret *= base
        if exponent < 0:
            ret = 1.0/ret
        return ret

    def randNumber(self):
        base = float(random.randint(-10, 10))
        exponent = random.randint(-10, 10)
        return base, exponent


if __name__ == "__main__":
    s = Solution()
    base, exponent = s.randNumber()
    print(base, exponent)
    ret = s.Power(base, exponent)
    print(ret)