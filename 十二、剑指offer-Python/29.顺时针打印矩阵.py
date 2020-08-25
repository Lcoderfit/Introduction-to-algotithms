"""
29.顺时针打印矩阵
时间复杂度：O(??)
空间复杂度：O(??)
"""
# -*- coding:utf-8 -*-
class Solution:
    # matrix类型为二维列表，需要返回列表
    def printMatrix(self, matrix):
        # write code here
        if not matrix or not matrix[0]:
            return None
        start = 0
        rows, cols = len(matrix), len(matrix[0])
        ret = list()
        while rows > start * 2 and cols > start * 2:
            tmp = self.printMatrixInCircle(matrix, start)
            ret.extend(tmp)
            start += 1

        return ret

    def printMatrixInCircle(self, matrix, start):
        rows = len(matrix)
        cols = len(matrix[0])
        endx = cols - start - 1
        endy = rows - start - 1

        tmp = list()

        for i in range(start, endx + 1):
            tmp.append(matrix[start][i])
        if start < endy:
            for i in range(start + 1, endy + 1):
                tmp.append(matrix[i][endx])
        if start < endx and start < endy:
            for i in range(endx - 1, start - 1, -1):
                tmp.append(matrix[endy][i])
        if start < endx and start < endy - 1:
            for i in range(endy - 1, start, -1):
                tmp.append(matrix[i][start])

        return tmp


if __name__ == "__main__":
    matrix = [
        [1, 2, 3, 4],
        [5, 6, 7, 8],
        [9, 10, 11, 12],
        [13, 14, 15, 16],
    ]

    s = Solution()
    ret = s.printMatrix(matrix)
    print(ret)