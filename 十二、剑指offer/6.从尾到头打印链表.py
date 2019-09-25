"""
6.从尾到头打印链表
时间复杂度：O(n)
空间复杂度：O(n)
"""
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def create_list(numbers):
    head = ListNode(numbers[0])
    p = head
    for i in range(1, len(numbers)):
        p.next = ListNode(numbers[i])
        p = p.next
    return head


class Solution:
    def printListFromTailToHead(self, listnode):
        p = listnode
        stack = list()
        while p != None:
            stack.append(p.val)
            p = p.next
        return stack[::-1]


if __name__ == "__main__":
    numbers = [6, 3, 1, 4, 2, 5]
    listnode = create_list(numbers)
    s = Solution()
    ret = s.printListFromTailToHead(listnode)
    print(ret)
