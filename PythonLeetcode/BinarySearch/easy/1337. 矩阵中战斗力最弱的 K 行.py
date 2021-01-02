"""

方法1： 二分查找
时间复杂度：O()
空间复杂度：O()

case1:
3 5 5
1 1 0 0 0
1 1 1 1 0
1 0 0 0 0
1 1 0 0 0
1 1 1 1 1

case2:
2 4 4
1 0 0 0
1 1 1 1
1 0 0 0
1 0 0 0

case3:
2 3 2
1 0
0 0
1 0


"""
import sys
from typing import List


class Solution:
    def k_weakest_rows(self, mat: List[List[int]], k: int) -> List[int]:
        m = len(mat)
        power = list()
        for i in range(m):
            power.append(self.find_rightmost(mat[i]))
        ans = list(range(m))
        # lambda j表示遍历ans列表元素, j对应遍历的ans元素值，首先按照power[j]来排序，如果有相同元素，则按照j来排序
        # 例如ans=[3, 1], power=[5, 2, 6, 7]，power最大索引必须要大于等于ans的最大元素（即最大索引必须>=3, 即长度>=4）
        # 首先lambda j即遍历ans得到3， 1； power[3] == 7, power[1] == 2，而power[3] < power[1]
        # 默认为升序排序，则1要排到3的前面，最终结果为ans=[1, 3]
        # (power[j], j)表示按照每行的战斗力从小到大排序，如果有两行战斗力相等，则行号小的排左边

        # Python提示object is not subscriptable的错误, 表示把不具有下标操作的对象a用成了a[i]
        # sorted返回一个列表，a.sort会改变a的值，但是返回值为None
        return sorted(ans, key=lambda j: (power[j], j))[:k]

    @staticmethod
    def find_rightmost(nums: List[int]) -> int:
        n = len(nums)
        if n == 0:
            return 0
        i, j = 0, n - 1
        while i <= j:
            mid = (i + j) // 2
            if nums[mid] == 1:
                i = mid + 1
            else:
                j = mid - 1
        return i

    @staticmethod
    def k_weakest_rows2(mat: List[List[int]], k: int) -> List[int]:
        n = len(mat)
        # sum函数
        power = [sum(line) for line in mat]
        ans = list(range(n))
        ans.sort(key=lambda i: (power[i], i))
        return ans[:k]


if __name__ == '__main__':
    s = Solution()
    k_cur, m_cur, n_cur = [int(i) for i in input().split(" ")]
    matrix = list()
    for i in range(m_cur):
        matrix.append([int(i) for i in input().split(" ")])
    res = s.k_weakest_rows(matrix, k_cur)
    print(res)
