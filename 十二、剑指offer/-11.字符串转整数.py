"""
-11.字符串转整数.py
时间复杂度：O(n)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
class Solution:
    def StrToInt(self, s):
        # write code here
        if not s:
            return 0
        t = s[0]
        if t == "-" or t == "+":
            s = s[1:]
        if not s.isdigit():
            return 0
        ret = 0
        for i in s:
            ret = ret*10 + (ord(i) - 48)
        p = 1 << 31
        q = -(1<<31) - 1
        if t == "-":
            ret = -ret
        return ret if ret < p and ret > q else 0


if __name__ == "__main__":
    s = input()
    st = Solution()
    ret = st.StrToInt(s)
    print(ret)
