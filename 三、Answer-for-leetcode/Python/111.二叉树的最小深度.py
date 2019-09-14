"""
104.二叉树的最大深度
时间复杂度：O()
空间复杂度：O()
"""
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None


def createTreeNode(node_list):
    if not node_list[0]:
        return None

    root = TreeNode(node_list)
    Nodes = [root]
    j = 1

    for node in Nodes:
        if node:
            node.left = (TreeNode(node_list[j]) if node_list[j] else None)
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
    def minDepth(self, root: TreeNode) -> int:
        if root:
            if root.right and root.left:
                return 1 + min(self.minDepth(root.right), self.minDepth(root.left))
            elif root.right:
                return 1 + self.minDepth(root.left)
            elif root.left:
                return 1 + self.minDepth(root.right)
            else:
                return 1
        else:
            return 0


if __name__ == "__main__":
    node_list = [3, 9, 20, None, None, 15, 7]
    root = createTreeNode(node_list)
    s = Solution()
    ret = s.minDepth(root)
    print(ret)