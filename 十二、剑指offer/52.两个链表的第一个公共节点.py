"""
52.两个链表的第一个公共节点.py
时间复杂度：O(n)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
class ListNode:
    def __init__(self, val):
        self.val = val
        self.next = None


def createLinkedList(listNode1, listNode2, listNode):
    if not listNode1 or not listNode2 or not listNode:
        return None

    p1 = pHead1 = ListNode(listNode1[0])
    p2 = pHead2 = ListNode(listNode2[0])

    for i in range(1, len(listNode1)):
        p1.next = ListNode(listNode1[i])
        p1 = p1.next

    for i in range(1, len(listNode2)):
        p2.next = ListNode(listNode2[i])
        p2 = p2.next

    p1.next = p2.next = ListNode(listNode[0])
    p = p1.next
    for i in range(1, len(listNode)):
        p.next = ListNode(listNode[i])
        p = p.next

    return pHead1, pHead2


class Solution:
    def FindFirstCommonNode(self, pHead1, pHead2):
        # write code here
        p1 = pHead1
        p2 = pHead2
        while p1 != p2:
            p1 = p1.next if p1 != None else pHead2
            p2 = p2.next if p2 != None else pHead1
        return p1

    def listOrder(self, pHead):
        p = pHead
        ret = []
        while p:
            ret.append(p.val)
            p = p.next
        return ret


if __name__ == "__main__":
    listNode1 = [1, 2, 3]
    listNode2 = [4, 5]
    listNode = [6, 7]
    pHead1, pHead2 = createLinkedList(listNode1, listNode2, listNode)

    s = Solution()
    ret1 = s.listOrder(pHead1)
    ret2 = s.listOrder(pHead2)
    print(ret1, ret2)

    ret = s.FindFirstCommonNode(pHead1, pHead2)
    print(ret.val)
