class BSTreeNode(object):
    def __init__(self, data):
        self.val = data
        self.leftChild = None
        self.rightChild = None


def creatTree(nodeList):
    if nodeList[0] == -1:
        return None
    head = BSTreeNode(nodeList[0])
    Nodes = [head]
    j = 1
    for node in Nodes:
        if node != None:
            node.leftChild = (BSTreeNode(nodeList[j]) if nodeList[j] != None else None)
            Nodes.append(node.leftChild)
            j += 1
            if j == len(nodeList):
                return head
            node.rightChild = (BSTreeNode(nodeList[j]) if nodeList[j] != None else None)
            j += 1
            Nodes.append(node.rightChild)
            if j == len(nodeList):
                return head


def levelOrder(root):
    if not root:
        return list()

    ret = []
    level = [root]
    while level:
        ret.append([node.val for node in level])

        temp = []
        for node in level:
            temp.extend([node.leftChild, node.rightChild])
        level = [leaf for leaf in temp if leaf]

    return ret



if __name__ == "__main__":
    node_list = [int(i) for i in input().split()]
    root = creatTree(node_list)
    ret = levelOrder(root)
    print(ret)