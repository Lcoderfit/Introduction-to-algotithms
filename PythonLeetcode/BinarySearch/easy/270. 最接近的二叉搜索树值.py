"""

方法1： 二分查找
时间复杂度：O()
空间复杂度：O()

case1:

"""
import sys


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    @staticmethod
    def closest_value(root: TreeNode, target: float) -> int:
        target = round(target)
        p = root
        q = None
        while p:
            q = p
            if p.val == target:
                return p.val
            elif p.val > target:
                p = p.left
            else:
                p = p.right
        return q.val


if __name__ == '__main__':
