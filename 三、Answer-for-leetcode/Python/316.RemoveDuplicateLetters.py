class Solution(object):
    def removeDuplicateLetters(self, s: str) -> str:
        dict = {}
        for c in s:
            if c not in dict:
                dict[c] = 1
            else:
                dict[c] += 1

        in_stack, stk = set(), []
        for c in s:
            if c not in in_stack:
                while stk and stk[-1] > c and dict[stk[-1]]:
                    in_stack.remove(stk.pop())
                stk += c
                in_stack.add(c)
            dict[c] -= 1
        return "".join(stk)


if __name__ == '__main__':
    while True:
        s = input()
        slu = Solution()
        res = slu.removeDuplicateLetters(s)
        print(res)