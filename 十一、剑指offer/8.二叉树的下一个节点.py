"""
8.二叉树的下一个节点
时间复杂度：O(logn)
空间复杂度：O(1)
"""


# -*- coding:utf-8 -*-
# class TreeLinkNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
#         self.next = None
class Solution:
    def GetNext(self, pNode):
        if not pNode:
            return None
        ret = None
        p = pNode
        if p.right:
            p = p.right
            while p.left:
                p = p.left
            ret = p
        elif p.next:
            while p.next and p.next.left != p:
                p = p.next
            ret = p.next

        return ret