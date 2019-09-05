# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def minDepth(self, root):
        # 有if root则需要判断左右孩子是否为空，如果最开始就判断了if not root,
        # 则后面可不用判断孩子节点是否为空
        if root:
            if root.left and root.right:
                return 1 + min(minDepth(root.left), minDepth(root.right))
            elif root.left:
                return 1 + min(minDepth(root.left))
            elif root.right:
                return 1 + min(minDepth(root.right))
            else:
                return 1
        else:
            return 0