"""

方法1： 二分查找
时间复杂度：O(zlogz)
空间复杂度：O(1)

case1:
5
case2:
9

方法2： 双指针
时间复杂度：O(z)
空间复杂度：O(1)

case1:
5
case2:
9
"""
import sys
from typing import List


class CustomFunction:
    @staticmethod
    def f(x: int, y: int) -> int:
        return x * y + 1


class Solution:
    @staticmethod
    def find_solution(custom_function: 'CustomFunction', z: int) -> List[List[int]]:
        res = []
        for i in range(1, 1001):
            if custom_function.f(i, 1) > z:
                break
            left, right = 1, 1000
            while left <= right:
                mid = (left + right) // 2
                if custom_function.f(i, mid) == z:
                    res.append([i, mid])
                    break
                # 变化方更大，如果是升序排序，就需要减小mid使得变化放变小
                elif custom_function.f(i, mid) > z:
                    right = mid - 1
                else:
                    left = mid + 1
        return res

    @staticmethod
    def find_solution2(custom_function: 'CustomFunction', z: int) -> List[List[int]]:
        i, j = 1, 1000
        res = list()
        while i <= 1000 and j >= 1:
            if custom_function.f(i, j) > z:
                j -= 1
            elif custom_function.f(i, j) < z:
                i += 1
            else:
                res.append([i, j])
                i += 1
                j -= 1
        return res


if __name__ == '__main__':
    s = Solution()
    custom_func_cur = CustomFunction()
    for line in sys.stdin:
        z_cur = int(line.strip())
        # print(s.find_solution(custom_func_cur, z_cur))
        print(s.find_solution2(custom_func_cur, z_cur))
