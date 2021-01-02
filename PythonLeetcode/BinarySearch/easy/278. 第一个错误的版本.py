"""

方法1： 二分查找
时间复杂度：O(logn)
空间复杂度：O(1)

case1:
5
"""
import sys
import random


# The isBadVersion API is already defined for you.
# @param version, an integer
# @return an integer
# def isBadVersion(version):


class Solution:
    def first_bad_version(self, n):
        """
        :type n: int
        :rtype: int
        """
        i, j = 1, n
        # 范围: [a, b]
        bad_version = random.randint(1, n)
        print("bad_version: ", bad_version)
        while i < j:
            # 一般情况下mid偏向的方向与指向mid的指针的方向相反
            mid = (i + j) // 2
            if not self.is_bad_version(mid, bad_version):
                i = mid + 1
            else:
                j = mid
        return i

    @staticmethod
    def is_bad_version(input_version, bad_version):
        return input_version >= bad_version


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        n_cur = int(line.strip())
        res = s.first_bad_version(n_cur)
        print(res)
