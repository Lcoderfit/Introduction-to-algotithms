'''
面试题37.序列化二叉树
时间复杂度：O()
空间复杂度：O()
'''
import json


class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None


class Codec:
    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """
        ret = list()
        if not root:
            return json.dumps(ret)
        level = [root]
        while any(level):
            for node in level:
                if node:
                    ret.append(node.val)
                else:
                    ret.append(None)

            tmp = list()
            for node in level:
                if node:
                    tmp.extend([node.left, node.right])
            level = tmp

        return json.dumps(ret)

    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """
        list_node = json.loads(data)
        if not list_node:
            return None
        root = TreeNode(list_node[0])
        if len(list_node) == 1:
            return root
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


if __name__ == "__main__":
    codec = Codec()
    data = "[1,2,3,null,null,4,5]"
    ret = codec.serialize(codec.deserialize(data))
    print(ret)

    data2 = "[1]"
    ret2 = codec.serialize(codec.deserialize(data2))
    print(ret2)
