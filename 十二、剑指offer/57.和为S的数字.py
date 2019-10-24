"""
56.数组中数字出现的次数.py
时间复杂度：O(n)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
class Solution:
    def FindNumbersWithSum(self, array, tsum):
        # write code here
        if not array:
            return None
        if len(array) <= 1:
            return None
        i, j = 0, len(array) - 1
        ret = None
        while i < j:
            if array[i] + array[j] == tsum:
                ret = (array[i], array[j])
                break
            elif array[i] + array[j] > tsum:
                j -= 1
            else:
                i += 1
        return ret


if __name__ == "__main__":
    array = [1, 2, 4, 7, 11, 15]
    tsum = 15
    s = Solution()
    ret = s.FindNumbersWithSum(array, tsum)
    print(ret)