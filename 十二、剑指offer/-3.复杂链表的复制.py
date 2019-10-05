"""
-3.复杂链表的复制
时间复杂度：O(n)
空间复杂度：O(n)
"""
# -*- coding:utf-8 -*-
class RandomListNode:
    def __init__(self, x):
        self.label = x
        self.next = None
        self.random = None


def CreateListNode(listNode1):
    if not listNode or not listNode2:
        return None
    if len(listNode1) != len(listNode2):
        return None

    pHead = RandomListNode(listNode[0])
    p = pHead
    for i in range(1, len(listNode)):
        p.next = RandomListNode(listNode[i])
        p = p.next


class Solution:
    # 返回 RandomListNode
    def Clone(self, pHead):
        # write code here
        if not pHead:
            return None
        p = pHead
        while p:
            q = RandomListNode(p.label)
            q.next = p.next
            p.next = q
            p = q.next

        p = pHead
        q = pHead.next
        while q:
            q.random = p.random.next
            p = p.next.next
            q = q.next.next

        p = pHead
        new_head = q = pHead.next
        while q:
            p.next = q.next
            p = q.next
            q.next = q.next.next
            q = p.next
        return new_head

    def ListOrder(self, pHead):
        if not pHead:
            return

        ret = list()
        p = pHead
        while p:
            ret.append(p.label)



if __name__ == "__main__":
    p = pHead = RandomListNode(1)
    




