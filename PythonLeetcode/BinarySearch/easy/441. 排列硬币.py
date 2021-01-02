"""

方法1： 二分查找
时间复杂度：O(logn)
空间复杂度：O(1)

case1:
5

case2:
8
"""
import sys


class Solution:
    @staticmethod
    def arrange_coins(n: int) -> int:
        i, j = 0, n
        # mid偏向的方向需要与指向mid的指针方向相反
        # 例如mid偏向右，则指向mid的指针为左指针，即i
        while i < j:
            mid = (i + j + 1) // 2
            cal = (mid*mid + mid) // 2
            if cal > n:
                j = mid - 1
            else:
                i = mid
        return i


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        n_cur = int(line.strip())
        res = s.arrange_coins(n_cur)
        print(res)
