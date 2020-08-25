"""
28.对称的二叉树.py
时间复杂度：O(n)
空间复杂度：O(1)
"""
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def createListNode(listNode):
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
    def isSymmetrical(self, pRoot):
        # write code here
        return self.isSymmetricalTree(pRoot, pRoot)

    def isSymmetricalTree(self, pRoot1, pRoot2):
        if not pRoot1 and not pRoot2:
            return True

        if not pRoot1 or not pRoot2:
            return False

        if pRoot1.val != pRoot2.val:
            return False

        return self.isSymmetricalTree(pRoot1.left, pRoot2.right) and \
               self.isSymmetricalTree(pRoot1.right, pRoot2.left)

    def levelOrder(self, root):
        if not root:
            return None
        level = [root]
        ret = list()
        while level:
            ret.append([node.val for node in level])

            tmp = list()
            for node in level:
                tmp.extend([node.left, node.right])
            level = [leaf for leaf in tmp if leaf]
        return ret


if __name__ == "__main__":
    listNode = [8, 6, 6, 5, 7, 7, 5]
    root = createListNode(listNode)

    s = Solution()
    ret = s.levelOrder(root)
    print(ret)

    isSym = s.isSymmetrical(root)
    print(isSym)


