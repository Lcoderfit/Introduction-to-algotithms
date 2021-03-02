"""

方法1： 模拟法
时间复杂度：O(logn)
空间复杂度：O(1)

方法2： 数学运算
时间复杂度：O(1)
空间复杂度：O(1)


case1:
16
r1:
True

case2:
5
r2:
False


(a + b) % p = (a % p + b % p) % p = (a % p + b) % p = (a + b % p) % p （1）
(a - b) % p = (a % p - b % p ) % p （2）
(a * b) % p = (a % p * b % p) % p （3）
a ^ b % p = ((a % p)^b) % p （4）
结合律：
((a+b) % p + c) % p = (a + (b+c) % p) % p （5）
((a*b) % p * c)% p = (a * (b*c) % p) % p （6）

交换律：
(a + b) % p = (b+a) % p （7）
(a * b) % p = (b * a) % p （8）

分配律：
(a+b) % p = ( a % p + b % p ) %p（9）
((a +b)% p * c) % p = ((a * c) % p + (b * c) % p) % p （10）

重要定理
若a≡b (% p)，则对于任意的c，都有(a + c)/ ≡ (b + c) (%p)；（11）
若a≡b (% p)，则对于任意的c，都有(a * c) ≡ (b * c) (%p)；（12）
若a≡b (% p)，c≡d (% p)，则 (a + c) ≡ (b + d) (%p)，(a - c) ≡ (b - d) (%p)，(a * c) ≡ (b * d) (%p)； （13）


(a+x)**k % a = x**k % a
4**k % 3 == (3+1)**k % 3 == 1**k % 3 == 1
注意，当(a+x)**k % a中的x为负数时，根据公式最终结果为 x**k % a, 负数取模运算规则如下：
-10 % 3 = -10 - （-10//3）*3 = -10 - (-4 * 3) = -10 + 12 = 2  #注意：-3.5向下取整为-4
10 % -3 = 10 - (10//3)*3 = 10 - 9 = 1
"""
import sys
from typing import List
from math import log2


class Solution:
    @staticmethod
    def is_power_of_four(n: int) -> bool:
        if n == 0:
            return False
        # 首先n % 4确定为4的倍数，然后在逐步除4即可
        while n % 4 == 0:
            n /= 4
        return n == 1

    # 设a=4**x, 则x = log4(a) = log2(a) / log2(4) = 1/2log2(a); 要使x为整数，则log2(a)必须为偶数
    @staticmethod
    def is_power_of_four2(n: int) -> int:
        # Python中浮点型对2取模也是可以正常计算的
        return (n > 0) and (log2(n) % 2 == 0)

    # 1.n & (n - 1)确保只有一位二进制为1
    # 2.由于为4的幂次方，所以其二进制表示中，为1的那一位一定在奇数位，例如100(4); 10000(16)
    @staticmethod
    def is_power_of_four3(n: int) -> int:
        return (n > 0) and (n & (n - 1) == 0) and (n & 0xaaaaaaaa == 0)

    @staticmethod
    def is_power_of_four4(n: int) -> int:
        # 4**k % 3 == (3+1)**k % 3 == 1**k % 3 == 1
        return (n > 0) and (n & (n - 1) == 0) and (n % 3 == 1)


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        n_cur = int(line)
        # print(s.is_power_of_four(n_cur))
        # print(s.is_power_of_four2(n_cur))
        print(s.is_power_of_four3(n_cur))
