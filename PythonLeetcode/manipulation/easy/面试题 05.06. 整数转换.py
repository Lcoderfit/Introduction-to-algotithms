"""

方法1： 布莱恩.克尼根算法
时间复杂度：O(1)
空间复杂度：O(1)

case1:
3 10
r1:
2

注意：位操作，本质上是补码与补码之间的操作，在Python中，整型，64位操作系统最少24个字节，32位最少12字节
64位：
    print(sys.getsizeof(0))         24
    print(sys.getsizeof(1))         28
    print(sys.getsizeof(2**30))     32
    print(sys.getsizeof(2**60))     36
    print(sys.getsizeof(2**90))     40
python整型内存扩容规律：从1开始，每增加30个二进制位则扩容4个字节

1.由于一个整型占用的内存数较其他语言比较大，所以例如-1，在Python中存储的形式为0b1111111111111111....1 (并不只有32个1，而是有8*28个1)；
所以Python中 (-1 ^ 1)其实是224位（24*8）的补码与224位的补码位运算
2.根据1可知，在Python中如果需要获取像其他语言整型的32位（4字节）表示，需要与0xFFFFFFFF相与取得32位补码（即原码取反得到反码，反码再+1得补码）
3.但是由于Python中整型占用内存的扩容机制，当x为负数时， while x: x = (x & (x - 1))就会一直循环下去，因为当x的值达到扩容阀值时，则会进行
扩容，负数扩容后，其补码扩容的bit位均为1，所以会有去不完的1给布莱恩克尼根算法操作。
4.如果需要查看Python中负数对应的32位补码，可以使用： bin(n & 0xFFFFFFFF)
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def convert_integer(a: int, b: int) -> int:
        # 获取整型的32位表示，由于Python整型占24个以上的字节，如果是负数，则其补码高位均为1，与0xFFFFFFFF相与后，高于32位的1均变为0
        # 就模拟出了整型的32位补码表示
        xor = (a & 0xFFFFFFFF) ^ (b & 0xFFFFFFFF)
        count = 0
        while xor:
            count += 1
            xor &= (xor - 1)
        return count

    @staticmethod
    def convert_integer2(a: int, b: int) -> int:
        return bin((a & 0xffffffff) ^ (b & 0xffffffff)).count('1')


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        a_cur, b_cur = [int(i) for i in line.split(" ")]
        print(s.convert_integer(a_cur, b_cur))
        # print(s.convert_integer2(a_cur, b_cur))
