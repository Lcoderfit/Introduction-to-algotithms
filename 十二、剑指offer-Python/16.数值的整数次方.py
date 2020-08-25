"""
16.数值的整数次方
时间复杂度：如果指数为n，则算法时间复杂度为O(n), 有一种高效解法为O(logn)
空间复杂度：O(1)
"""
import random


# -*- coding:utf-8 -*-
class Solution1:
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


class Solution2:
    def Power(self, base, exponent):
        if base == 0.0 and exponent <= 0:
            return 0.0
        t_exponent = abs(int(exponent))
        ret = self.PowerCore(base, t_exponent)
        if exponent < 0:
            ret = 1.0/ret
        return ret

    def PowerCore(self, base, exponent):
        if exponent & 0xffffffff == 0:
            return 1
        if exponent == 1:
            return base

        ret = self.PowerCore(base, exponent>>1)
        ret *= ret
        if exponent & 1 == 1:
            ret = ret * base
        return ret


    def randNumber(self):
        base = float(random.randint(-10, 10))
        exponent = random.randint(-10, 10)
        return base, exponent



if __name__ == "__main__":
    s = Solution2()
    base, exponent = s.randNumber()
    print("base: %s" % base)
    print("exponent: %s" % exponent)
    ret = s.Power(base, exponent)
    print(ret)