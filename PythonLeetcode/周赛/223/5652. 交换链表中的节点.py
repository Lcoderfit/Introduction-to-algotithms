"""

方法1： 双指针
时间复杂度：O(n)
空间复杂度：O(1)

输入：head = [1,2,3,4,5], k = 2
输出：[1,4,3,2,5]
示例 2：

输入：head = [7,9,6,6,7,8,3,0,9,5], k = 5
输出：[7,9,6,6,8,7,3,0,9,5]
示例 3：

输入：head = [1], k = 1
输出：[1]
示例 4：

输入：head = [1,2], k = 1
输出：[2,1]
示例 5：

输入：head = [1,2,3], k = 2
输出：[1,2,3]
"""
import sys
from typing import List


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    @staticmethod
    def swap_nodes(head: ListNode, k: int) -> ListNode:
        p = head
        link_len = 0
        while p:
            link_len += 1
            p = p.next

        distance = link_len + 1 - 2 * k
        font_steps = abs(distance)
        p = head
        for i in range(font_steps):
            p = p.next

        steps = k if distance >= 0 else k - font_steps
        steps -= 1
        q = head
        if steps > 0:
            for i in range(steps):
                p = p.next
                q = q.next
        p.val, q.val = q.val, p.val
        return head


if __name__ == '__main__':
    pass
