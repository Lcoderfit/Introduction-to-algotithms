"""
24.反转链表
时间复杂度：O(n)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def createListNode(list_node):
    if not list_node:
        return None
    head = ListNode(list_node[0])
    p = head
    for i in range(1, len(list_node)):
        p.next = ListNode(list_node[i])
        p = p.next
    return head


class Solution:
    # 返回ListNode
    def ReverseList1(self, pHead):
        # write code here
        if (not pHead) or (not pHead.next):
            return pHead
        p, q = pHead, pHead.next
        if not q.next:
            q.next = p
            p.next = None
            return q
        r = q.next
        pHead.next = None
        while r:
            q.next = p
            p = q
            q = r
            r = r.next
        q.next = p
        return q

    def ReverseList2(self, pHead):
        if not pHead:
            return pHead
        p, q, r = None, None, pHead
        while r.next !=None:
            p = q
            q = r
            r = r.next
            q.next = p
        r.next = q

        return r

    def printList(self, pHead):
        ret = list()
        p = pHead
        while p:
            ret.append(p.val)
            p = p.next
        return ret


if __name__ == "__main__":
    list_node = [1, 4, 6, 9, 3, 2]
    pHead = createListNode(list_node)
    s = Solution()
    print(s.printList(pHead))

    reverse_head = s.ReverseList(pHead)
    ret = s.printList(reverse_head)
    print(ret)
