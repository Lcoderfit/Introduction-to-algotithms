"""
542.01矩阵-01-matrix
类型：BFS
"""
from typing import List


class Solution:
    def updateMatrix(self, matrix: List[List[int]]) -> List[List[int]]:
        row = len(matrix)
        col = len(matrix[0])
        for i in range(row):
            for j in range(col):
                l, u = 10001, 10001
                if matrix[i][j] != 0:
                    if i > 0:
                        u = matrix[i - 1][j]
                    if j > 0:
                        l = matrix[i][j - 1]
                    matrix[i][j] = min(l, u) + 1
        for i in range(row -1, -1, -1):
            for j in range(col - 1, -1, -1):
                r, d = 10001, 10001
                if matrix[i][j] != 0:
                    if i < row - 1:
                        d = matrix[i + 1][j]
                    if j < col - 1:
                        r = matrix[i][j + 1]
                    matrix[i][j] = min(matrix[i][j], min(r, d) + 1)

        return matrix