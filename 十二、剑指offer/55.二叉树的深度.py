"""
55.二叉树的深度.py
时间复杂度：O(n)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None


def CreateTreeNode(listNode):
    if not listNode:
        return None
    root = TreeNode(listNode[0])
    level = [root]
    j = 1
    for node in level:
        if node:
            node.left = (TreeNode(listNode[j]) if listNode[j] != None else None)
            level.append(node.left)
            j += 1
            if j == len(listNode):
                return root

            node.right = (TreeNode(listNode[j]) if listNode[j] != None else None)
            level.append(node.right)
            j += 1
            if j == len(listNode):
                return root


class Solution:
    def TreeDepth(self, pRoot):
        # write code here
        if not pRoot:
            return 0
        return 1 + max(self.TreeDepth(pRoot.left), self.TreeDepth(pRoot.right))

    def LevelOrder(self, pRoot):
        if not pRoot:
            return list()
        ret = list()
        level = [pRoot]

        while level:
            ret.append([node.val for node in level])

            tmp = list()
            for node in level:
                tmp.extend([node.left, node.right])
            level = [leaf for leaf in tmp if leaf]
        return ret


if __name__ == "__main__":
    listNode = [1, 2, 3, 4, 5, None, 6, None, None, 7, None, None, None]
    pRoot = CreateTreeNode(listNode)
    s = Solution()
    level_ret = s.LevelOrder(pRoot)
    print(level_ret)
    ret = s.TreeDepth(pRoot)
    print(ret)
