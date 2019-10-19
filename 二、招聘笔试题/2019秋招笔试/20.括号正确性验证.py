class Solution:
    def Validate(self, string):
        if not string:
            return 0
        stack = list()
        stack.append(string[0])
        for i in range(1, len(string)):
            if stack and self.Check(stack[-1], string[i]):
                stack.pop()
            else:
                stack.append(string[i])
        ret = 0 if stack else 1
        return ret

    def Check(self, s1, s2):
        if (s1 == '(' and s2 == ')') or (s1 == '<' and s2 == '>') \
        or (s1 == '[' and s2 == ']') or (s1 == '{' and s2 == '}'):
            return True
        return False


if __name__ == "__main__":
    string = input().strip()
    s = Solution()
    ret = s.Validate(string)
    print(ret)