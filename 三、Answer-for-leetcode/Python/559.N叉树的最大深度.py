"""
559.N叉树的最大深度
时间复杂度：O(n)
空间复杂度：O(1)
"""
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.children = children


class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        if not root:
            return 0
        children_depth = [self.maxDepth(node) for node in root.children]
        if children_depth:
            return max(children_depth) + 1
        return 1
