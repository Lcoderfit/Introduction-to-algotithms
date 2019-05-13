class Solution(object):
    def reverse_bits(self, n: int) -> int:
        #题目要求返回一个整数
        #return bin(n)[-1:1:-1].ljust(32, '0')
        return int(bin(n)[-1:1:-1].ljust(32, '0'), 2)


if __name__ == '__main__':
    while True:
        tmp = '0b' + input()
        n = int(tmp, 2)
        s = Solution()
        res = s.reverse_bits(n)
        print(res)
