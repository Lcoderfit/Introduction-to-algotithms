'''
面试题68 - I. 二叉搜索树的最近公共祖先.py
时间复杂度：O()
空间复杂度：O()
'''


# 迭代解法
# 时间: O(n): 最坏情况，一条单链表，只有左孩子
# 空间: O(1)
class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if not root:
            return None
        if p.val == q.val:
            return p
        while root:
            if root.val < q.val and root.val < p.val:
                root = root.right
            if root.val > q.val and root.val > p.val:
                root = root.left
            else:
                return root


# 递归解法
# 时间: O(n): 最坏情况，一条单链表，只有左孩子
# 空间: O(n): 递归栈空间
class Solution1:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if not root:
            return None
        if p.val > q.val:
            p, q = q, p
        if p.val == root.val or q.val == root.val:
            return root
        if p.val < root.val and q.val > root.val:
            return root
        if p.val < root.val and q.val < root.val:
            return self.lowestCommonAncestor(root.left, p, q)
        if p.val > root.val and q.val > root.val:
            return self.lowestCommonAncestor(root.right, p, q)