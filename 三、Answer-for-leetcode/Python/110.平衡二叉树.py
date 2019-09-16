"""
104.二叉树的最大深度
时间复杂度：O(n)
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

    length = len(node_list)
    root = TreeNode(node_list[0])
    Nodes = [root]
    j = 1

    for node in Nodes:
        if node:
            node.left = (TreeNode(node_list[j]) if node_list[j] else None)
            Nodes.append(node.left)
            j += 1
            if length == j:
                return root

            node.right = (TreeNode(node_list[j]) if node_list[j] else None)
            Nodes.append(node.right)
            j += 1
            if length == j:
                return root


class Solution:
    def is_balance(self, root: TreeNode) -> bool:
        return self.depth(root) != -1

    def depth(self, root: TreeNode) -> int:
        if not root:
            return 0
        left_depth = self.depth(root.left)
        if left_depth == -1:
            return -1

        right_depth = self.depth(root.right)
        if right_depth == -1:
            return -1

        return max(left_depth, right_depth) + 1 if abs(left_depth - right_depth) < 2 else -1


if __name__ == "__main__":
    node_list = [3, 9, 20, None, None, 15, 7]
    root = createTreeNode(node_list)
    s = Solution()
    ret = s.is_balance(root)
    print(ret)