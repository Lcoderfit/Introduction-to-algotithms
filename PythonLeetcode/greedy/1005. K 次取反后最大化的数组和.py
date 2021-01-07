"""

方法1： 排序+贪心
时间复杂度：O(nlogn)
空间复杂度：O(1)

方法1： 优化排序+贪心
时间复杂度：O(nlogn)
空间复杂度：O(1)

case1:
4 2 3
1
r:
5

case2:
3 -1 0 2
3
r:
6

case3:
2 -3 -1 5 -4
2
r:
13

case4:
4 -5 4 -5 9 4 5
1
r:
26

通过绝对值排序后，即使将绝对值更大的负数转换为相反数之后，排序顺序也不会变
例如：-3 1 -1 4 -5 3   绝对值排序后： 1 -1 -3 3 4 -5
从右往左遍历，将负数变为其相反数， =》 1 1 3 3 4 5, 排序顺序也不会变，且a[0]一直为最小值
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def largest_sum_after_k_negations(a: List[int], k: int) -> int:
        a.sort()
        for i in range(len(a)):
            if a[i] >= 0 or k == 0:
                break
            k -= 1
            a[i] = -a[i]
        if k % 2 == 0:
            return sum(a)
        a.sort()
        a[0] = -a[0]
        return sum(a)

    @staticmethod
    def largest_sum_after_k_negations2(a: List[int], k: int) -> int:
        # 精妙之处
        a.sort(key=lambda x: abs(x))
        # 按照绝对值排序之后（从小到大排序），从右边开始扫描, 即使经过取相反数的操作后排序顺序也不会变
        for i in range(len(a) - 1, -1, -1):
            if (k > 0) and (a[i] < 0):
                k -= 1
                a[i] *= -1
        if k & 1:
            a[0] *= -1
        return sum(a)


if __name__ == '__main__':
    s = Solution()
    a_cur = [int(i) for i in input().strip().split(" ")]
    k_cur = int(input().strip())
    # print(s.largest_sum_after_k_negations(a_cur, k_cur))
    print(s.largest_sum_after_k_negations2(a_cur, k_cur))
