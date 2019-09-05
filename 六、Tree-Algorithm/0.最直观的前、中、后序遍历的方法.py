# class Solution:
#     def inorderTraversal(self, root: TreeNode) -> List[int]:
#         WHITE, GRAY = 0, 1
#         res = []
#         stack = [(WHITE, root)]
#         while stack:
#             color, node = stack.pop()
#             if node is None: continue
#             if color == WHITE:
#                 stack.append((WHITE, node.right))
#                 stack.append((GRAY, node))
#                 stack.append((WHITE, node.left))
#             else:
#                 res.append(node.val)
#         return res

from typing import List


class Solution1:
    """
    颜色标记法-中旭遍历
    """
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        WHITE, GRAY = 0, 1
        ret = []
        stack = [(WHITE, root)]
        while stack:
            color, node = stack.pop()
            if not node: continue
            if color == WHITE:
                stack.append((WHITE, node.right))
                stack.append((GRAY, node))
                stack.append((WHITE, node.left))
            else:
                ret.append(node.val)
        return ret


class Solution2:
    """
    颜色标记法——先序遍历
    """
    def preorderTraversal(self, root: TreeNode) -> List[int]:
        WHITE, GRAY = 0, 1
        ret = []
        stack = [(WHITE, root)]

        while stack:
            color, node = stack.pop()
            if not node: continue
            if color == WHITE:
                stack.append((WHITE, node.right))
                stack.append((WHITE, node.left))
                stack.append((GRAY, node))
            else:
                ret.append(node.val)

        return ret


class Solution3:
    """
    颜色标记法——后序遍历
    """
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        WHITE, GRAY = 0, 1
        stack = [(WHITE, root)]
        ret = []
        while stack:
            color, node = stack.pop()
            if not node: continue
            if color == WHITE:
                stack.append((GRAY, node))
                stack.append((WHITE, node.right))
                stack.append((WHITE, node.left))
            else:
                ret.append(node.val)

        return ret



