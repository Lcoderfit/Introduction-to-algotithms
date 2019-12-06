class Node:
    def __init__(self, key=None, color=None, size=0):
        self.key = key
        self.color = color
        self.size = size
        self.left = None
        self.right = None
        self.p = None


class Tree:
    def __init__(self):
        self.root = None
        self.nil = Node()


class Solution:
    def leftRotate(self, T: Tree, x: Node):
        y = x.right
        x.right = y.left
        if y.left != T.nil:
            y.left.p = x
        y.p = x.p
        if x.p == T.nil:
            T.root = y
        elif x.p.left == x:
            x.p.left = y
        else:
            x.p.right = y
        y.left = x
        x.p = y

        y.size = x.size
        x.size = x.left.size + x.right.size + 1

    def rightRotate(self, T: Tree, x: Node):
        y = x.left
        x.left = y.right
        if y.right != T.nil:
            y.right.p = x
        y.p = x.p
        if x.p == T.nil:
            T.root = y
        elif x.p.left == x:
            x.p.left = y
        else:
            x.p.right = y
        y.right = x
        x.p = y

        y.size = x.size
        x.size = x.left.size + x.right.size + 1

    def