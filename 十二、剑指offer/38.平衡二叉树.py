"""
38.平衡二叉树
时间复杂度：O(n)
空间复杂度：O(n)
"""
class Solution:
    def IsBalanced_Solution(self, pRoot):
        # write code here
        if not pRoot:
            return True
        left = self.depth(pRoot.left)
        right = self.depth(pRoot.right)
        if abs(left - right) < 2:
            return True
        return self.IsBalanced_Solution(pRoot.left) and self.IsBalanced_Solution(pRoot.right)

    def depth(self, pRoot):
        if not pRoot:
            return 0
        left = self.depth(pRoot.left)
        right = self.depth(pRoot.right)

        return max(left, right) + 1