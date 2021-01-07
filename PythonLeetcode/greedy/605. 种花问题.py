"""

方法1： 模拟
时间复杂度：O(n)
空间复杂度：O(1)

case1:
1 0 0 0 1
1
r:
True

case2:
1 0 0 0 1
2
r:
False

case3:
0
1
r:
True
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def can_plan_flowers(flowerbed: List[int], n: int) -> bool:
        if n == 0:
            return True
        i = 0
        while i < len(flowerbed):
            if flowerbed[i] == 1:
                i += 2
                continue

            if i + 1 < len(flowerbed):
                if flowerbed[i + 1] == 1:
                    i += 2
                elif i > 0 and flowerbed[i - 1] == 1:
                    i += 1
                else:
                    flowerbed[i] = 1
                    n -= 1
                    i += 2
            else:
                if (i == 0) or (i > 0 and flowerbed[i - 1] == 0):
                    n -= 1
                i += 1

            if n == 0:
                return True
        return False

    @staticmethod
    def can_plan_flowers2(flowerbed: List[int], n: int) -> bool:
        # 跳格子
        i, length = 0, len(flowerbed)
        count = 0
        while i < length:
            if flowerbed[i] == 1:
                i += 2
            elif i == length - 1 or flowerbed[i + 1] != 1:
                count += 1
                i += 2
            else:
                i += 3
        return count >= n


if __name__ == '__main__':
    s = Solution()
    flowerbed_cur = [int(i) for i in input().strip().split(" ")]
    n_cur = int(input().strip())
    print(s.can_plan_flowers(flowerbed_cur, n_cur))
