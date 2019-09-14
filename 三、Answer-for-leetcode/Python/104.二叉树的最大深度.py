"""
104.二叉树的最大深度
时间复杂度：O(logn)
空间复杂度：O(1)
"""
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.right = None
        self.left = None


def createTreeNode(node_list):
    if not node_list[0]:
        return None

    root = TreeNode(node_list[0])
    Nodes = [root]
    j = 1

    for node in Nodes:
        if node:
            node.left = (TreeNode(node_list[j]) if  node_list[j] else None)
            Nodes.append(node.left)
            j += 1
            if len(node_list) == j:
                return root

            node.right = (TreeNode(node_list[j]) if node_list[j] else None)
            Nodes.append(node.right)
            j += 1
            if len(node_list) == j:
                return root


class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        if not root:
            return 0
        return 1 + max(self.maxDepth(root.left), self.maxDepth(root.right))


if __name__ == "__main__":
    node_list = [3, 9, 20, None, None, 15, 7]
    root = createTreeNode(node_list)

    s = Solution()
    ret = s.maxDepth(root)
    print(ret)
