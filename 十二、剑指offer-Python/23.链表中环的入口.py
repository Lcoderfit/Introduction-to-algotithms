"""
23.链表中环的入口.py
时间复杂度：O(n)
空间复杂度：O(1)
"""
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def createListNode(l1, l2):
    if (not l1) or (not l2):
        return None
    p = pHead = ListNode(l1[0])
    for i in range(1, len(l1)):
        p.next = ListNode(l1[i])
        p = p.next
    t = p.next = ListNode(l2[0])
    p = p.next
    for i in range(1, len(l2)):
        p.next = ListNode(l2[i])
        p = p.next
    p.next = t

    return pHead



class Solution:
    def EntryNodeOfLoop1(self, pHead):
        if (not pHead) or (not pHead.next):
            return None

        p = pHead.next
        q = pHead.next.next
        while p != q:
            p = p.next
            q = q.next.next
        r = pHead
        while r != p:
            r = r.next
            p = p.next
        return p

    def EntryNodeOfLoop2(self, pHead):
        pFast = pSlow = pHead
        while pFast and pFast.next:
            pFast = pFast.next.next
            pSlow = pSlow.next

            if pFast == pSlow:
                break

        if (not pFast) or (not pFast.next):
            return None

        pFast = pHead
        while pFast != pSlow:
            pFast = pFast.next
            pSlow = pSlow.next
        return pFast

    def printListNode(self, pHead):
        p = pHead
        ret = list()
        while p:
            ret.append(p.val)
            p = p.next


if __name__ == "__main__":
    l1 = [1, 2]
    l2 = [3, 4, 5]
    pHead = createListNode(l1, l2)
    s = Solution()
    ret = s.EntryNodeOfLoop2(pHead)
    print(ret.val)
