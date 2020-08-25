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
    @staticmethod
    def print_list_from_tail_to_head(node_list=None):
        stack = list()
        while node_list is not None:
            stack.append(node_list.val)
            node_list = node_list.next
        return stack[::-1]


if __name__ == "__main__":
    array = [6, 3, 1, 4, 2, 5]
    link_list = create_list(array)
    s = Solution()
    ret = s.print_list_from_tail_to_head(link_list)
    print(ret)
