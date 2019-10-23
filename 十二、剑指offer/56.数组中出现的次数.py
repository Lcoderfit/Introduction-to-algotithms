"""
56.数组中出现的次数.py
时间复杂度：O(N)
空间复杂度：O(1)
0与非0数异或==这个数
"""
# -*- coding:utf-8 -*-
class Solution:
    # 返回[a,b] 其中ab是出现一次的两个数字
    def FindNumsAppearOnce(self, array):
        # write code here
        if not array or len(array) <= 1:
            return None
        tmp = 0
        for i in array:
            tmp ^= i
        pivot = tmp - (tmp & (tmp - 1))
        one, two = 0, 0
        for i in array:
            if pivot & i:
                one ^= i
            else:
                two ^= i
        return (one, two)


if __name__ == "__main__":
    array = [int(i) for i in input().split(" ")]
    s = Solution()
    ret = s.FindNumsAppearOnce(array)
    print(ret)
