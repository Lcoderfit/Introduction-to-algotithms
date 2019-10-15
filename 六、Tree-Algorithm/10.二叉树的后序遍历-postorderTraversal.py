from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        """
        递归法
        时间复杂度：O(n)，每个节点恰好访问一次
        空间复杂度：O(n)，最坏情况下是每个节点都存储一次
        """
        if not root:
            return list()
        return self.postorderTraversal(root.left) + self.postorderTraversal(root.right) + [root]


class Solution1:
    """
    迭代法，用两个栈实现
    时间复杂度：O(n)，每个节点恰好访问一次
    空间复杂度：O(n)，最坏情况下是每个节点都存储一次
    """
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        if not root:
            return list()
        ret, stack = [], [root]
        while stack:
            node = stack.pop()
            ret.append(node.val)
            if node.left:
                stack.append(node.left)
            if node.right:
                stack.append(node.right)

        # 列表逆置后的元素顺序就相当于栈元素的出栈顺序
        return ret[::-1]
