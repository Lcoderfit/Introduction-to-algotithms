"""
时间限制：C/C++语言 1000MS；其他语言 3000MS
内存限制：C/C++语言 65536KB；其他语言 589824KB
题目描述：
将一个 32 位有符号整数，每位上的数字进行反转，正负号保持不变。

输入
一个32位有符号整数

输出
对于每个测试实例，输出其反转后的数字。如果反转后整数溢出那么就返回 0
"""
import sys


class Solution:
    def reverse_integer(self, n):
        ret = n if n > 0 else abs(n)
        ret = int(str(ret)[::-1])
        if n >= 0 and (ret > 0x7fffffff):
            return 0
        if n < 0 and (ret > 0xffffffff):
            return 0
        if n < 0:
            ret  = -ret
        return ret


if __name__ == "__main__":
    s = Solution()
    for line in sys.stdin:
        n = int(line.strip())
        ret = s.reverse_integer(n)
        print(ret)
    