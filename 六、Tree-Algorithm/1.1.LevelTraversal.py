"""
二叉树层序遍历：宽度优先遍历
时间复杂度：
空间复杂度：
"""


class BinaryTravelsal:
    def levelOrder(self, root):
        """
        二叉树层序遍历, 用队列的方法实现
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if not root:
            return []
        ret = []
        # 存储每一层的元素
        level = [root]
        while level:
            ret.append([node.val for node in level])
            # 存储每层中每个元素的孩子节点
            temp = []
            for node in level:
                temp.extend([node.left, node.right])
            level = [leaf for leaf in temp if leaf]
        return ret
