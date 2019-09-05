"""
二叉树的先序遍历
"""
from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    """
    递归法
    时间复杂度：O(n)，每个节点恰好访问一次
    空间复杂度：O(n)，最坏情况下是每个节点都存储一次
    """
    def preorderTraversal(self, root: TreeNode) -> List[int]:
        if not root:
            return list()

        return [root.val] + self.preorderTraversal(root.left) + self.preorderTraversal(root.right)


class Solution1:
    """
    迭代法, 用栈来实现
    时间复杂度：O(n)，每个节点恰好访问一次
    空间复杂度：O(n)，最坏情况下是每个节点都存储一次
    """
    def preorderTraversal(self, root: TreeNode) -> List[int]:
        if not root:
            return list()

        ret = []
        stack = [root]
        while stack:
            cur = stack.pop()
            ret.append(cur.val)

            if cur.right:
                stack.append(cur.right)
            if cur.left:
                stack.append(cur.left)

        return ret



