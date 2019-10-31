"""
-4.字符串的排列
时间复杂度：O(n!)
空间复杂度：O(???)
"""
# -*- coding:utf-8 -*-
class Solution:
    def Permutation(self, ss):
        # write code here
        if not ss:
            return list()
        if len(ss) <= 1:
            return [ss]
        ret = list()
        for i in range(len(ss)):
            t = ss[:i] + ss[i+1:]
            t_list = self.Permutation(t)
            for j in t_list:
                ret.append(ss[i] + j)
        return sorted(list(set(ret)))


if __name__ == "__main__":
    ss = "abc"
    s = Solution()
    ret = s.Permutation(ss)
    print(ret)
