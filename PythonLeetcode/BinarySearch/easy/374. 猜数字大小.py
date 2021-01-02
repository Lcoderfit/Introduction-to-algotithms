"""

方法1： 二分查找
时间复杂度：O(logn)
空间复杂度：O(1)

case1:
10
guess: 6
"""
import sys
import random


# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
# def guess(num: int) -> int:

class Solution:
    def guess_number(self, n: int) -> int:
        i, j = 1, n
        generate_number = random.randint(1, n)
        print("generate_number: ", generate_number)
        # 循环终止状态，i最终会指向目标位置, j固定偏向左边，i固定偏向右边
        while i <= j:
            mid = (i + j) // 2
            if self.guess(generate_number, mid) > 0:
                i = mid + 1
            else:
                j = mid - 1
        return i

    @staticmethod
    def guess(generate_number, input_number):
        if generate_number > input_number:
            return 1
        if generate_number < input_number:
            return -1
        return 0


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        n_cur = int(line.strip())
        res = s.guess_number(n_cur)
        print(res)
