# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return list()

        ret = []
        level = [root]

        order = 0
        while level:
            copy_level = []
            if order % 2:
                copy_level = level[::-1]
                order = 0
            else:
                copy_level = level
                order = 1
            level_val = [node.val for node in copy_level]
            ret.append(level_val)

            temp = []
            for node in level:
                temp.extend([node.left, node.right])
            level = [leaf for leaf in temp if leaf]

        return ret


# # Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
#
#
# class Solution:
#     def zigzagLevelOrder(self, root):
#         res = []
#         def helper(root, depth):
#             if not root: return
#             if len(res) == depth:
#                 res.append([])
#             if depth % 2 == 0:
#                 res[depth].append(root.val)
#             else:
#                 res[depth].insert(0, root.val)
#             helper(root.left, depth + 1)
#             helper(root.right, depth + 1)
#
#         helper(root, 0)
#         return res