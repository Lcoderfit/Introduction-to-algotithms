"""
-6.按Z字形顺序打印二叉树.py
时间复杂度：O(n)
空间复杂度：O(n)
"""
# -*- coding:utf-8 -*-
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def createTreeNode(listNode):
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
    def Print(self, pRoot):
        # write code here
        if not pRoot:
            return []
        level = [pRoot]
        count = 1
        ret = []
        while level:
            t = [node.val for node in level]
            if count:
                ret.append(t)
                count = 0
            else:
                ret.append(t[::-1])
                count = 1
            tmp = []
            for node in level:
                tmp.extend([node.left, node.right])
            level = [leaf for leaf in tmp if leaf]
        return ret

    def levelOrder(self, root):
        if not root:
            return None
        level = [root]
        ret = []
        while level:
            ret.append([node.val for node in level])

            tmp = []
            for node in level:
                tmp.extend([node.left, node.right])
            level = [leaf for leaf in tmp if leaf]
        return ret


if __name__ == "__main__":
    listNode = [8,6,10,5,7,9,11]
    root = createTreeNode(listNode)

    s = Solution()
    level_ret = s.levelOrder(root)
    print(level_ret)

    ret = s.Print(root)
    print(ret)
