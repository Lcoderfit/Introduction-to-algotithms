# """
# # Definition for a Node.
# class Node:
#     def __init__(self, val, children):
#         self.val = val
#         self.children = children
# """
# class Solution:
#     def levelOrder(self, root: 'Node') -> List[List[int]]:
#         ret = []
#         if not root:
#             return []
#         nlev, netlev = [root], []
#         while nlev:
#             re = []
#             for n in nlev:
#                 re.append(n.val)
#                 for i in n.children:
#                     netlev.append(i)
#             nlev, netlev = netlev, []
#             ret.append(re)
#         return ret

"""
# Definition for a Node.
"""
class Node:
    def __init__(self, val, children):
        self.val = val
        self.children = children


class Solution:
    def levelOrder(self, root: Node) -> List[List[int]]:
        if not root:
            return list()

        ret = []
        level = [root]

        while level:
            ret.append([node.val for node in level])

            temp = []
            for node in level:
                temp.extend(node.children)
            level = [leaf for leaf in temp if leaf]
        return ret