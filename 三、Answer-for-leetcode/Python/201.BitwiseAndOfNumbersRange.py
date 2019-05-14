# class Solution:
#     def rangeBitwiseAnd(self, m: int, n: int) -> int:
#         # n&(n-1)会把最后一个1后面所有位都置为0,有点类似找m和n二进制的公共前缀
#         count = 0  # 记录右移的次数
#         while m != n:  # 直到m和n相同
#             m >>= 1
#             n >>= 1
#             count += 1
#         return m<<count


class Solution1:
    def range_bitwise_and(self, m: int, n: int) -> int:
        count = 0
        while m != n:
            m >>= 1
            n >>= 1
            count += 1
        return m<<count


class Solution2:
    '''???'''
    def range_bitwise_and(self, m: int, n: int) -> int:
        while n > m:
            n &= (n-1)
        return n


if __name__ == '__main__':
    while True:
        m, n = [int(i) for i in input().split()]
        s = Solution()
        res = s.range_bitwise_and(m, n)
        print(res)