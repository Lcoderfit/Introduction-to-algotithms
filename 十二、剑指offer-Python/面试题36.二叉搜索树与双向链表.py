'''
面试题36. 二叉搜索树与双向链表
时间复杂度：
空间复杂度：
'''

# class Solution:
#     def treeToDoublyList(self, root: 'Node') -> 'Node':
#         if not root:
#             return root
#         first, last = None, None
#         def dfs(head:'Node'):
#             nonlocal first, last
#             if head:
#                 dfs(head.left)
#                 if last:
#                    last.right = head
#                    head.left = last
#                 else:
#                     first = head
#                 last = head
#                 dfs(head.right)
#         dfs(root)
#         last.right = first
#         first.left = last
#         return first

class TreeNode:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def tree_to_duply_list(self, root: TreeNode) -> TreeNode:
        if not root:
            return None

        stack = list()

        def helper(node: TreeNode):
            if not node:
                return
            if node.left:
                helper(node.left)

            if stack:
                stack[-1].right = node
            stack.append(node)
            if len(stack) >= 2:
                stack[-1].left = stack[-2]

            if node.right:
                helper(node.right)
        helper(root)
        stack[0].left = stack[-1]
        stack[-1].right = stack[0]

        return stack[0]


if __name__ == "__main__":
    print("Lcoderfit")