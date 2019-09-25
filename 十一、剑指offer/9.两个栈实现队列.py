# -*- coding:utf-8 -*-
"""
9.两个栈实现队列
时间复杂度：push: O(1)  pop: O(n)
空间复杂度：O(n)
"""
class Solution:
    def __init__(self):
        self.stack1 = list()
        self.stack2 = list()
    def push(self, node):
        # write code here
        self.stack1.append(node)
    def pop(self):
        # return xx
        if (not self.stack1) and (not self.stack2):
            return -1
        elif not self.stack2:
            while self.stack1:
                p = self.stack1.pop()
                self.stack2.append(p)
        ret = self.stack2.pop()
        return ret


if __name__ == "__main__":
    s = Solution()
    for i in range(10):
        s.push(i)
    for i in range(10):
        print(s.pop())