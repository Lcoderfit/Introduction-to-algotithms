"""
-1.二叉搜索树的后序遍历序列
时间复杂度：O(n)
空间复杂度：O(n)
"""
#-*- coding:utf-8 -*-
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
    def VerifySquenceOfBST(self, sequence):
        if not sequence:
            return False
        return self.isBinary(sequence)

    def isBinary(self, sequence):
        if len(sequence) <= 1:
            return True

        ret = False
        pivot = sequence[-1]
        i = 0
        while (i < len(sequence) - 1) and sequence[i] < pivot:
            i += 1
        mid = i
        for i in range(mid, len(sequence) - 1):
            if sequence[i] < pivot:
                return False
        return self.isBinary(sequence[0:mid]) and self.isBinary(sequence[mid:-1])

    def levelOrder(self, root):
        if not root:
            return root
        level = [root]
        ret = []
        while level:
            ret.append([node.val for node in level])

            tmp = list()
            for node in level:
                tmp.extend([node.left, node.right])
            level = [leaf for leaf in tmp if leaf]
        return ret


if __name__ == "__main__":
    listNode = [5, 3, 8, 2, 4, 7, 10]
    root = createTreeNode(listNode)

    s = Solution()
    level_ret = s.levelOrder(root)
    print(level_ret)

    sequence = [2, 4, 3, 7, 10, 8, 5]
    ret = s.VerifySquenceOfBST(sequence)
    print(ret)

