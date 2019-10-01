"""
26.数的子结构
时间复杂度：O(??)
空间复杂度：O(??)
"""
# -*- coding:utf-8 -*-
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

def createTreeNode(list_node):
    if not list_node:
        return None
    root = TreeNode(list_node[0])
    level = [root]
    j = 1
    for node in level:
        if node:
            node.left = (TreeNode(list_node[j]) if list_node[j] != None else None)
            level.append(node.left)
            j += 1
            if j == len(list_node):
                return root

            node.right = (TreeNode(list_node[j]) if list_node[j] != None else None)
            level.append(node.right)
            j += 1
            if j == len(list_node):
                return root



class Solution:
    def HasSubtree(self, pRoot1, pRoot2):
        if (not pRoot1) or (not pRoot2):
            return False

        res = self.isSubtree(pRoot1, pRoot2)
        if not res:
            res = self.HasSubtree(pRoot1.left, pRoot2)
        if not res:
            res = self.HasSubtree(pRoot1.right, pRoot2)
        return res

    def isSubtree(self, pRoot1, pRoot2):
        if not pRoot2:
            return True
        if not pRoot1:
            return False

        if not self.Equal(pRoot1.val, pRoot2.val):
            return False

        return self.isSubtree(pRoot1.left, pRoot2.left) and self.isSubtree(pRoot1.right, pRoot2.right)

    def Equal(self, num1, num2):
        return abs(num1 - num2) < 0.0000001

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
    list_node1 = [8, 8, 7, 9, 2, None, None, None, None, 4, 7]
    list_node2 = [8, 9, 2]

    pRoot1 = createTreeNode(list_node1)
    pRoot2 = createTreeNode(list_node2)

    s = Solution()
    ret1 = s.levelOrder(pRoot1)
    ret2 = s.levelOrder(pRoot2)
    print(ret1)
    print(ret2)
    ret = s.HasSubtree(pRoot1, pRoot2)
    print(ret)

