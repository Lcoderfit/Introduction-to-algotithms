"""
31.栈的压入弹出顺序.py
时间复杂度：O(n)
空间复杂度：O(n)
"""
class Solution:
    def IsPopOrder(self, pushV, popV):
        if not pushV or not popV:
            return False
        if len(pushV) == 0:
            return False

        stack = list()
        j = 0
        for i in range(len(pushV)):
            stack.append(pushV[i])
            while stack and stack[-1] == popV[j]:
                stack.pop()
                j += 1
        return stack == []


if __name__ == "__main__":
    pushV = [1, 2, 3, 4, 5]
    popV = [4, 5, 3, 2, 1]
    s = Solution()
    ret = s.IsPopOrder(pushV, popV)
    print(ret)