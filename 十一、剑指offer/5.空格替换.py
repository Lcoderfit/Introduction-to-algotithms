"""
5.空格替换
时间复杂度：O(n)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
class Solution:
    # s 源字符串
    def replaceSpace(self, s):
        # write code here
        str_list = list(s)
        ret = ""
        print(str_list)
        for i in range(len(str_list)):
            ret += str_list[i] if str_list[i] != " " else "%20"
        return ret

if __name__ == "__main__":
    string = "hello world"
    s = Solution()
    ret = s.replaceSpace(string)
    print(ret)
