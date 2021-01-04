"""

方法1： 位运算
时间复杂度：O(n)
空间复杂度：O(1)

case1:
255.0.0.7 10
r:

"""
import sys
from typing import List


class Solution:
    def ip_to_cidr(self, ip, n):
        ans = []
        start = self.ip_to_int(ip)
        # TODO: 0.0.0.0貌似不可分。。。。。
        if start == 0:
            ans.append("0.0.0.0/32")
            n -= 1
            start += 1

        while n:
            # start & -start算出来的是start的二进制表示中，最右边的一个“1”及该“1”右边的所有0
            # 用于构建子网的位必须都为0，否则不能用于构建本题中的ip
            # 能够用于构建子网的位不能比n的二进制表示的长度要大，n二进制表示为100，则用于构建子网的位必须 < 3
            # 而由于start & (-start）除了包含用于构建子网的“0”，还包含start二进制表示中最右边的一个1，
            # 故(start & -start).bit_length() <= n.bit_length(), 所以在符合要求情况下，最长的用于构建子网的位为:
            # min((start & -start).bit_length(), n.bit_length()) - 1
            # mask = 32 - （min((start & -start).bit_length(), n.bit_length()) - 1） =》
            # 33 - min((start & -start).bit_length(), n.bit_length()
            mask = 33 - min((start & -start).bit_length(), n.bit_length())
            ans.append(self.int_to_ip(start) + "/" + str(mask))
            # 例如用了3位构建子网，也就是说start末尾为1000, 则下次应该从第四位+1的数值开始，也就是 1 << (32 - mask)
            start += 1 << (32 - mask)
            n -= 1 << (32 - mask)
        return ans

    @staticmethod
    def ip_to_int(ip):
        ans = 0
        for x in ip.split("."):
            ans = ans * 256 + int(x)
        return ans

    @staticmethod
    def int_to_ip(x):
        # (x >> i) % 256 跟10进制数a 使用 a%10取个位数一个道理
        return ".".join(str((x >> i) % 256) for i in (24, 16, 8, 0))


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        params = [i for i in line.strip()]
        ip_cur, n_cur = params[0], int(params[1])
        print(s.ip_to_cidr(ip_cur, n_cur))
