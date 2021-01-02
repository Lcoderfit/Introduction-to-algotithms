"""

方法1： 二分查找
时间复杂度：O()
空间复杂度：O()

case1:

"""
import sys


class Solution:
    @staticmethod
    def is_perfect_square(num: int) -> bool:
        i, j = 0, num
        while i < j:
            mid = (i + j) // 2
            mul = mid*mid
            if mul == num:
                return True
            elif mul > num:
                j = mid - 1
            else:
                i = mid + 1
        if i*i != num:
            return False
        return True

    @staticmethod
    def is_perfect_square1(num: int) -> bool:
        i, j = 1, num
        while i < j:
            mid = (i + j) // 2
            mul = mid * mid
            if mul < num:
                i = mid + 1
            else:
                j = mid
        return i*i == num


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        num_cur = int(line.strip())
        # res = s.is_perfect_square(num_cur)
        res = s.is_perfect_square1(num_cur)
        print(res)
