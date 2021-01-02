"""

方法1： 
时间复杂度：O(D)
空间复杂度：O(1)

case1:
6 10
r1:
4

case2:
990 1048
r1:
28

1.L, R 是 L <= R 且在 [1, 10^6] 中的整数。
2.R - L 的最大值为 10000。
注意：计算置位就是数字二进制表示中为1的位的总数
10**3=1000  2**10=1024  所以：10**3 < 2**10
由上可以推到：
10**6 < 2**20
10**9 < 2**30

由于L,R位于[1, 10^6区间]所以其二进制表示最多也不超过20位，最大的质数计算置位为19
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def count_prime_set_bits(l: int, r: int) -> int:
        primes = {2, 3, 5, 7, 11, 13, 17, 19}
        # sum函数会将bool类型转换为0和1求和，例如sum[False, True],结果相当于sum([0, 1])
        return sum(bin(i).count("1") in primes for i in range(l, r + 1))


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        l_cur, r_cur = [int(i) for i in line.split(" ")]
        res = s.count_prime_set_bits(l_cur, r_cur)
        print(res)
