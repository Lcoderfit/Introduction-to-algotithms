"""
-10.左旋字符串.py
时间复杂度：O(n)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
class Solution:
    def LeftRotateString(self, s, n):
        # write code here
        if not s:
            return s
        t = n % len(s)
        return s[t:] + s[0:t]


if __name__ == "__main__":
    string = "abcXYZdef"
    n = 3
    s = Solution()
    ret = s.LeftRotateString(string, n)
    print(ret)

