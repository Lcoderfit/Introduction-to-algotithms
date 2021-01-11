"""

方法1： 线性扫描（线性探测）
时间复杂度：O(n)
空间复杂度：O(1)

case1:
2 1
r:
false

case2:
3 5 5
r:
false

case3:
0 3 2 1
r:
true
"""
import sys
from typing import List


# 这个不好在一个循环中处理的逻辑，可以分成两步，第一步找到峰顶，第二步从峰顶向右，不满足条件则返回
class Solution:
    @staticmethod
    def valid_mountain_array(arr: List[int]) -> bool:
        i = 0
        # 递增扫描
        while i + 1 < len(arr) and arr[i] < arr[i + 1]:
            i += 1

        # 峰顶元素不能为第一个或者最后一个
        if (i == 0) or (i == len(arr) - 1):
            return False

        # 递减扫描
        while i + 1 < len(arr) and arr[i] > arr[i + 1]:
            i += 1

        return i == len(arr) - 1


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        arr_cur = [int(i) for i in line.strip().split(" ")]
        print(s.valid_mountain_array(arr_cur))
