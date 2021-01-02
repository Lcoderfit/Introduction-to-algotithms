"""
1064. 不动点.py
方法1： 二分查找
时间复杂度：O(logn)
空间复杂度：O(1)

case1:
0 1 2 3 4

case2:
-10 -5 0 3 7
"""
import sys
from typing import List


class Solution:
    def fixedPoint(self, A: List[int]) -> int:
        if not A:
            return -1
        i, j = 0, len(A) - 1
        # 找定值用：<, 这样循环结束后就是i==j, 即指向最终目标
        # 找临界值用: <=, 循环结束后就是i == (j - 1), 指向目标的不是i就是j，看实际情况取
        while i < j:
            # 因为当j指向符合要求的值时，可能还有更小的值，所以另mid偏向i，且当mid指向符合条件的结果时，
            # j直接指向mid（因为有可能是mid，也有可能是比mid小的位置）
            mid = (i + j) // 2
            if A[mid] >= mid:
                j = mid
            elif A[mid] < mid:
                i = mid + 1
        if A[i] == i:
            return i
        return -1


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        nums = [int(i) for i in line.strip().split(" ")]
        res = s.fixedPoint(nums)
        print(res)
