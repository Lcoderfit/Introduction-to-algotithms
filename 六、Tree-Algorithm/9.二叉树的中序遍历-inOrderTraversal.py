from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        """
        递归法
        时间复杂度：O(n)，每个节点恰好访问一次
        空间复杂度：O(n)，最坏情况下是每个节点都存储一次
        """
        if not root:
            return list()

        return self.inorderTraversal(root.left) + [root] + self.inorderTraversal(root.right)


class Solution1:
    """
   非 递归法
    时间复杂度：O(n)，每个节点恰好访问一次
    空间复杂度：O(n)，最坏情况下是每个节点都存储一次
    """
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        ret, stack, p = [], [], root
        while p or stack:
            while p:
                stack.append(p)
                p = p.left

            cur = stack.pop()
            ret.append(cur.val)
            p = cur.right

        return ret


