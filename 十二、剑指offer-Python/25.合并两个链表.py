"""
-2.合并两个链表
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
    head = ListNode(list_node[0])
    p = head
    for i in range(1, len(list_node)):
        p.next = ListNode(list_node[i])
        p = p.next
    return head


class Solution:
    # 返回合并后列表
    def Merge(self, pHead1, pHead2):
        # write code here
        p = merge = ListNode(0)
        while pHead1 and pHead2:
            if pHead1.val <= pHead2.val:
                merge.next = pHead1
                pHead1 = pHead1.next
            else:
                merge.next = pHead2
                pHead2 = pHead2.next
            merge = merge.next
        merge.next = pHead1 or pHead2

        return p.next

    def printList(self, head):
        p = head
        ret = list()
        while p:
            ret.append(p.val)
            p = p.next
        return ret


if __name__ == "__main__":
    list_node1 = [1, 3, 4, 6, 12]
    list_node2 = [2, 7, 9, 10, 13]
    head1 = createListNode(list_node1)
    head2 = createListNode(list_node2)
    s = Solution()
    print(s.printList(head1))
    print(s.printList(head2))
    m = s.Merge(head1, head2)
    ret = s.printList(m)
    print(ret)