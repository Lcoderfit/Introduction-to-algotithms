"""
21.调整数组顺序使奇数位于偶数前面
时间复杂度：O(n)
空间复杂度：O(n)
"""
# -*- coding:utf-8 -*-
# 常规法，要求奇数与奇数，偶数与偶数的相对位置不变
# 时间复杂度：O(n^2)
class Solution1:
    def reOrderArray(self, array):
        # write code here
        ret = list()
        index = 0
        for i in range(len(array)):
            if array[i] % 2:
                ret.insert(index, array[i])
                index += 1
            else:
                ret.append(array[i])
        return ret


# 相对位置可以变, 算法时间复杂度: O(n)
class Solution2:
    def reOrderArray(self, array):
        # write code here
        i, j = 0, len(array) - 1
        while i < j:
            while ((array[i] & 1) == 1):
                i += 1
            while ((array[j] & 1) == 0):
                j -= 1

            if i < j:
                array[i], array[j] = array[j], array[i]
                i += 1
                j -= 1
            else:
                break
        return array


if __name__ == "__main__":
    array = [1, 3, 6, 1, 2, 5, 7]
    s = Solution2()
    ret = s.reOrderArray(array)
    print(ret)