"""
39.数组中出现次数超过一半的数字
时间复杂度：O(n)
空间复杂度：O(1)
"""
# -*- coding:utf-8 -*-
class Solution:
    def MoreThanHalfNum_Solution(self, numbers):
        # write code here
        if not numbers:
            return 0
        ret = numbers[0]
        count = 1
        for i in range(1, len(numbers)):
            if count == 0:
                ret = numbers[i]
                count = 1
            if numbers[i] == ret:
                count += 1
            else:
                count -= 1
        times = 0
        for i in range(len(numbers)):
            if numbers[i] == ret:
                times += 1
        if times * 2 <= len(numbers):
            ret = 0
        return ret


if __name__ == "__main__":
    numbers = [1, 2, 3, 2, 2, 2, 5, 4, 2]

    s = Solution()
    ret = s.MoreThanHalfNum_Solution(numbers)
    print(ret)