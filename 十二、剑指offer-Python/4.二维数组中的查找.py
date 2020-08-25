"""
4.二维数组中的查找
时间复杂度：O(m + n)
空间复杂度：O(1)
"""
class Solution:
    def find(self, target, array):
        n, m = len(array), len(array[0])
        i, j = 0, m - 1
        while (i >= 0 and i < n) and (j >= 0 and j < m):
            if array[i][j] == target:
                return True
            elif array[i][j] > target:
                j -= 1
            else:
                i += 1
        return False


if __name__ == "__main__":
    array = [
        [1, 2, 8, 9],
        [2, 4, 9, 12],
        [4, 7, 10, 13],
        [6, 8, 11, 15],
    ]
    target = 7

    s = Solution()
    ret = s.find(target, array)
    print(ret)
