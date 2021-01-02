"""

方法1： 二分查找
时间复杂度：O(logn)
空间复杂度：O(1)

方法2： 线性扫描
时间复杂度：O(n)
空间复杂度：O(1)

case1:
a
c f j

case 2:
c
c f j

case 3:
d
c f j

case 4:
g
c f j

case 5:
j
c f j

case 6:
k
c f j
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def next_greatest_letter(letters: List[str], target: str) -> str:
        i, j = 0, len(letters) - 1
        # 本质上是求左边界，因为是求满足比目标值大的数中的最小值，在升序的数组里，“最小”对应的就是左边界
        # 左边界，mid指向左，右指针指向mid
        while i < j:
            mid = (i + j) // 2
            if letters[mid] <= target:
                i = mid + 1
            else:
                j = mid
        if (i == len(letters) - 1) and (letters[i] > target):
            return letters[-1]
        if (i == len(letters) - 1) and (letters[i] <= target):
            return letters[0]
        return letters[i]

    @staticmethod
    def next_greatest_letter1(letters: List[str], target: str) -> str:
        i, j = 0, len(letters) - 1
        while i <= j:
            mid = (i + j) // 2
            if letters[mid] <= target:
                i = mid + 1
            else:
                j = mid - 1
        if i == len(letters):
            return letters[0]
        return letters[i]

    @staticmethod
    def next_greatest_letter2(letters: List[str], target: str) -> str:
        for c in letters:
            if c > target:
                return c
        return letters[0]


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        target_cur = line.strip()
        letters_cur = [i for i in input().split(" ")]
        res = s.next_greatest_letter(letters_cur, target_cur)
        print(res)
