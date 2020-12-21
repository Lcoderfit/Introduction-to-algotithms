"""
1351. 统计有序矩阵中的负数.py
方法1： 二分查找
时间复杂度：O(mlogn)
空间复杂度：O(1)

法2：倒序遍历
时间复杂度：O(m)
空间复杂度：O(1)

case:
4 4
4 3 2 -1
3 2 1 -1
1 1 -1 -2
-1 -1 -2 -3
"""
import sys
from typing import List


class Solution:
    def count_negatives(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        res = 0
        pos = 0
        for i in range(m - 1, -1, -1):
            pos = self.binary_search_negatives(grid[i], pos)
            res += n - pos
        return res

    @staticmethod
    def binary_search_negatives(nums: List[int], pos: int) -> int:
        i, j = pos, len(nums) - 1
        while i <= j:
            mid = (i + j) // 2
            if nums[mid] < 0:
                j = mid - 1
            else:
                i = mid + 1
        return i

    @staticmethod
    def count_negatives2(grid: List[List[int]]) -> int:
        pos = 0
        m, n = len(grid), len(grid[0])
        res = 0
        for i in range(m - 1, -1, -1):
            for j in range(pos, n):
                if grid[i][j] < 0:
                    pos = j
                    res += n - pos
                    break
        return res


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        m_cur, n_cur = [int(i) for i in line.strip("").split(" ")]
        grid_cur = []
        for i in range(m_cur):
            grid_cur.append([int(i) for i in input().strip().split(" ")])
        # res_cur = s.countNegatives(grid_cur)
        res_cur = s.count_negatives2(grid_cur)
        print(res_cur)
