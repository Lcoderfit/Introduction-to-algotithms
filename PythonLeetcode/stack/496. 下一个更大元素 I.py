"""

方法1： 单调栈（单调递减栈，当当前元素大于栈顶元素时，则一直出栈直到栈顶元素大于当前元素，如此可以构建一个从栈底到栈顶单调递减的栈）
时间复杂度：O(m+n)
空间复杂度：O(n)

case1:
4 1 2
1 3 4 2
r:
-1 3 -1

case2:
2 4
1 2 3 4
r:
3 -1

单调栈的特性：一般用于求元素右边或者左边第一个比自己大或者小的数
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def next_greater_element(nums1: List[int], nums2: List[int]) -> List[int]:
        hash_table = {}
        stack = []
        for i in range(len(nums2)):
            # 如果当前元素大于栈顶元素，则将栈顶元素弹出直到当前元素小于栈顶元素为止
            while stack and (nums2[i] > stack[-1]):
                hash_table[stack.pop()] = nums2[i]
            stack.append(nums2[i])

        while stack:
            hash_table[stack.pop()] = -1
        return [hash_table[v] for v in nums1]


if __name__ == '__main__':
    s = Solution()
    nums1_cur = [int(i) for i in input().strip().split(" ")]
    nums2_cur = [int(i) for i in input().strip().split(" ")]
    print(s.next_greater_element(nums1_cur, nums2_cur))
