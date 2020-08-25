"""
-1.链表中倒数第k个节点
时间复杂度：O(n)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def createListNode(list_node):
    if not list_node:
        return None
    pHead = ListNode(list_node[0])
    p = pHead
    for i in range(1, len(list_node)):
        p.next = ListNode(list_node[i])
        p = p.next

    return pHead


class Solution:
    def FindKthToTail(self, head, k):
        # write code here
        if not head:
            return None
        p, q = head, head
        count = 1
        while q.next and (count < k):
            count += 1
            q = q.next
        if count != k:
            return None
        while q.next:
            p = p.next
            q = q.next
        return p

    def printList(self, head):
        p = head
        ret = list()
        while p:
            ret.append(p.val)
            p = p.next
        return ret


if __name__ == "__main__":
    list_node = [1, 4, 5, 8, 3, 2]
    head = createListNode(list_node)

    s = Solution()
    ret = s.printList(head)
    print(ret)
    p = s.FindKthToTail(head, 5)
    print(p.val)
