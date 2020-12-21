import sys


class Solution:
    @staticmethod
    def my_sqrt(x: int) -> int:
        i, j = 1, x
        while i < j:
            mid = int((i + j + 1) / 2)
            if mid * mid > x:
                j = mid - 1
            else:
                i = mid
        return j


if __name__ == '__main__':
    s = Solution()
    # array = [int(i) for i in input().split()]
    param = int(input().strip())
    print(s.my_sqrt(param))
