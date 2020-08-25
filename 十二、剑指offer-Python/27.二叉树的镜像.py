"""
27.二叉树的镜像
时间复杂度：O(n)
空间复杂度：O(1)
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
            node.left = (TreeNode(listNode[j] if listNode[j] != None else None))
            level.append(node.left)
            j += 1
            if j == len(listNode):
                return root

            node.right = (TreeNode(listNode[j] if listNode[j] != None else None))
            level.append(node.right)
            j += 1
            if j == len(listNode):
                return root


class Solution:
    # 返回镜像树的根节点
    def Mirror(self, root):
        # write code here
        if not root:
            return
        root.left, root.right = root.right, root.left
        if root.left:
            self.Mirror(root.left)
        if root.right:
            self.Mirror(root.right)

    def levelOrder(self, root):
        if not root:
            return None
        level = [root]
        ret  = list()
        while level:
            ret.append([node.val for node in level])

            tmp = list()
            for node in level:
                tmp.extend([node.left, node.right])
            level = [leaf for leaf in tmp if leaf]

        return ret


if __name__ == "__main__":
    listNode = [8, 6, 10, 5, 7, 9, 11]
    root = createTreeNode(listNode)

    s = Solution()
    old_ret = s.levelOrder(root)
    print(old_ret)

    s.Mirror(root)
    n_ret = s.levelOrder(root)
    print(n_ret)
