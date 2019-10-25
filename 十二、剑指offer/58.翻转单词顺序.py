"""
58.翻转单词顺序.py
时间复杂度：O(n)
空间复杂度：O(n)
"""
# -*- coding:utf-8 -*-
class Solution:
    def ReverseSentence(self, s):
        # write code here
        return " ".join(s.split(" ")[::-1])


if __name__ == "__main__":
    string = "I am a student."
    s = Solution()
    ret = s.ReverseSentence(string)
    print(ret)