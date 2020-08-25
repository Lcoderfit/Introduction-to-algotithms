"""
-3.从上往下打印二叉树
时间复杂度：O(n)
空间复杂度：O(n)
"""
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

# 可改进，参考递归创建方法？？？？
def createTree(nodeList):
    if nodeList[0] == None:
        return None
    root = TreeNode(nodeList[0])
    level = [root]
    j = 1
    for node in level:
        if node != None:
            node.left = (TreeNode(nodeList[j]) if nodeList[j] != None else None)
            level.append(node.left)
            j += 1
            if j == len(nodeList):
                return root
            node.right = (TreeNode(nodeList[j]) if nodeList[j] != None else None)
            level.append(node.right)
            j += 1
            if j == len(nodeList):
                return root


class Solution:
    # 返回从上到下每个节点值列表，例：[1,2,3]
    def PrintFromTopToBottom(self, root):
        # write code here
        ret = list()
        if not root:
            return ret
        level = [root]
        while level:
            ret.extend([node.val for node in level])

            tmp = []
            for node in level:
                tmp.extend([node.left, node.right])
            level = [leaf for leaf in tmp if leaf]

        return ret


if __name__ == "__main__":
    nodeList = [1, 2, 4, None, 5, 6, None, 7, None, None, None]
    root = createTree(nodeList)
    s = Solution()
    ret = s.PrintFromTopToBottom(root)
    print(ret)

