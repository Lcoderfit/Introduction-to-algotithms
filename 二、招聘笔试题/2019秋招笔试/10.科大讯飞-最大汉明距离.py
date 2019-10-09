class Solution:
    def max_length(self, string):
        if not string:
            return 0
        string = string.split(";")
        ret = 0
        for s in string:
            if s.isalpha():
                ret = ret + abs((ord(s[0]) - ord(s[-1])))
        return ret


if __name__ == "__main__":
    string = "aasdkgfj;askdf;bkdsjfe;pzdjakdk"

    s = Solution()
    ret = s.max_length(string)
    print(ret)
