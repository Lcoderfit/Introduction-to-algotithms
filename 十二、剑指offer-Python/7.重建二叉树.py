"""
7.重建二叉树
时间复杂度：O(n^2) 最坏情况下（左子节点相连的单链），遍历inorder切片，n + (n-1) + (n-2) ... 1 = n*(n+1)/2
空间复杂度：O(n)
"""
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def reConstructBinaryTree(self, pre, tin):
        if (not pre) or (not tin):
            return
        root = TreeNode(pre[0])
        index = -1
        for i in range(len(tin)):
            if tin[i] == pre[0]:
                index = i
                break

        if index == -1:
            return
        l_tin, r_tin = tin[:index], tin[index + 1:]
        length = len(l_tin)
        l_pre, r_pre = pre[1:length + 1], pre[length + 1:]
        root.left = self.reConstructBinaryTree(l_pre, l_tin)
        root.right = self.reConstructBinaryTree(r_pre, r_tin)

        return root

    def preOrder(self, root):
        if not root:
            return list()
        level = [root]
        ret = []
        while any(level):
            t = []
            for node in level:
                t.append(node.val if node else None)
            ret.append(t)

            tmp = []
            for node in level:
                if node:
                    tmp.extend([node.left, node.right])

            level = [leaf for leaf in tmp]

        return ret

    def creatTree(self, nodeList):
        if nodeList[0] == None:
            return None
        root = TreeNode(nodeList[0])
        Nodes = [root]
        j = 1
        for node in Nodes:
            if node != None:
                node.left = (TreeNode(nodeList[j]) if nodeList[j] != None else None)
                Nodes.append(node.left)
                j += 1
                if j == len(nodeList):
                    return root
                node.right = (TreeNode(nodeList[j]) if nodeList[j] != None else None)
                j += 1
                Nodes.append(node.right)
                if j == len(nodeList):
                    return root


if __name__ == "__main__":
    pre = [1, 2, 4, 7, 3, 5, 6, 8]
    tin = [4, 7, 2, 1, 5, 3, 8, 6]

    level = [1, 2, 3, 4, None, 5, 6, None, 7, None, None, 8, None]

    s = Solution()
    root = s.reConstructBinaryTree(pre, tin)
    # root = s.creatTree(level)
    ret = s.preOrder(root)
    print(ret)

