"""
236.二叉树的公共祖先
时间复杂度：O(n)
空间复杂度：O(n)
"""
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.right = None
        self.left = left


def creatTree(nodeList):
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


class Solution:
    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        dic = {root: None}
        def bfs(node):
            if node.left:
                dic[node.left] = node
            if node.right:
                dic[node.right] = node
            bfs(node.left)
            bfs(node.right)

        bfs(root)
        l1, l2 = p, q
        while l1 != l2:
            l1 = dic.get(l1) if l1 else q
            l2 = dic.get(l2) if l2 else p

        return l1


if __name__ == "__main__":
    node_list = []
    array = input().split()
    for i in array:
        if i.isdigit():
            node_list.append(int(i))
        else:
            node_list.append(None)

    p, q = [TreeNode(int(val)) for val in input().split()]

    root = creatTree(node_list)
    s = Solution()
    ret = s.lowestCommonAncestor(root, p, q)
    print(ret)


