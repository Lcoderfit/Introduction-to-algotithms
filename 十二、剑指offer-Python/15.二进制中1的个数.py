"""
15.二进制中1的个数
时间复杂度：如果n表示数字大小，那么时间复杂度：O(logn)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
class Solution:
    def NumberOf1(self, n):
        # write code here
        count = 0
        while n & 0xffffffff != 0:  # Python的整型没有位数限制，所以需要限制位数，当n<0时会无限循环
                                    # 当n<0时，例如ob1000000000000000(最前面一个1表示负数)，n & (n - 1)会在后面无限补0
            count += 1
            n = (n - 1) & n
        return count


if __name__ == "__main__":
    s = Solution()
    ret = s.NumberOf1(15)
    print(ret)