"""
12.矩阵中的路径
时间复杂度：???
空间复杂度：????
"""
# -*- coding:utf-8 -*-
class Solution:
    # def hasPath(self, matrix, rows, cols, path):
    #     # write code here
    #     if (not matrix) or (not path) \
    #         or (rows < 1) or (cols < 1):
    #         return False
    #     plen = [0]
    #     visited = [[0] * cols] * rows
    #     for i in range(rows):
    #         for j in range(cols):
    #             if i == 0 and j == 1:
    #                 t = []
    #                 print("Lcoderfit")
    #                 self.findPath(matrix, rows, cols, path, i, j, plen, visited, t)
    #                 print(t)
    #                 print("Lcoderfit")
    #             if self.findPath(matrix, rows, cols, path, i, j, plen, visited):
    #                 return True
    #     return False
    #
    # def findPath(self, matrix, rows, cols, path, i, j, plen, visited, t = None):
    #     if plen[0] == len(path):
    #         return True
    #     print(plen, (i, j))
    #     has_path = False
    #     if (i >= 0 and i < rows) and (j >= 0 and j < cols) and \
    #         (path[plen[0]] == matrix[i][j]) and (not visited[i][j]):
    #         print("h1")
    #         print(matrix[i][j])
    #         if t != None:
    #             print("fuck")
    #             t.append(matrix[i][j])
    #         plen[0] += 1
    #         visited[i][j] = True
    #         has_path = self.findPath(matrix, rows, cols, path, i + 1, j, plen, visited, t) \
    #                 or self.findPath(matrix, rows, cols, path, i, j + 1, plen, visited, t) \
    #                 or self.findPath(matrix, rows, cols, path, i - 1, j, plen, visited, t) \
    #                 or self.findPath(matrix, rows, cols, path, i, j - 1, plen, visited, t)
    #
    #         if not has_path:
    #             plen[0] -= 1
    #             visited[i][j] = False
    #
    #     return has_path


    def hasPath(self, matrix, rows, cols, string):
        if (not matrix) or (rows < 1) or (cols < 1) or (not string):
            return False
        visited = [[0]*cols for _ in range(rows)]
        pathLength = 0
        for row in range(rows):
            for col in range(cols):
                if self.hasPathCore(matrix, rows, cols, row, col, string, pathLength, visited):
                    return True
        return False

    def hasPathCore(self, matrix, rows, cols, row, col, string, pathLength, visited):
        if len(string) == pathLength:
            return True
        hasPath = False
        if row >= 0 and row < rows and col >= 0 and col < cols and matrix[row][col] == string[pathLength] and (not visited[row][col]):
            pathLength += 1
            visited[row][col] = 1

            hasPath = self.hasPathCore(matrix, rows, cols, row, col - 1, string, pathLength, visited) \
                    or self.hasPathCore(matrix, rows, cols, row - 1, col, string, pathLength, visited) \
                    or self.hasPathCore(matrix, rows, cols, row, col + 1, string, pathLength, visited) \
                    or self.hasPathCore(matrix, rows, cols, row + 1, col, string, pathLength, visited)

            if not hasPath:
                pathLength -= 1
                visited[row][col] = 0

        return hasPath


    # 针对matrix为字符的情况
    # def hasPath(self, matrix, rows, cols, path):
    #     # 参数校验
    #     if len(matrix) == 0 or len(matrix) != rows * cols or len(path) == 0:
    #         return False
    #     visited = [False] * len(matrix)
    #     pathlength = [0]
    #     for i in range(rows):
    #         for j in range(cols):
    #             # 以矩阵中每一个位置作为起点进行搜索
    #             if self.haspath(matrix, rows, cols, path, j, i, visited, pathlength):
    #                 return True
    #     return False
    #
    # # 判断矩阵位置（x,y）的字符能否加入已找到的路径中
    # def haspath(self, matrix, rows, cols, path, x, y, visited, pathlength):
    #     '''
    #     :param matrix:字符矩阵
    #     :param rows:矩阵的行数
    #     :param cols:矩阵的列数
    #     :param path:需要寻找的路径
    #     :param x:当前位置的横坐标(对应列数)
    #     :param y:当前位置的纵坐标(对应行数)
    #     :param visited:访问标志数组
    #     :param pathlength:已经找到的路径长度
    #     :return:是否存在路径
    #     '''
    #     if pathlength[0] == len(path):
    #         return True
    #     curhaspath = False
    #     # 参数校验：1、位置坐标不超过行列数 2、当前位置字符等于路径中对应位置的字符 3、当前位置未存在于当前已找到的路径中
    #     if 0 <= x < cols and 0 <= y < rows \
    #             and matrix[y * cols + x] == path[pathlength[0]] \
    #             and not visited[y * cols + x]:
    #
    #         visited[y * cols + x] = True
    #         pathlength[0] += 1
    #         # 分别向左，向右，向上，向下移动一个格子，任一方向能够继续往下走均可
    #         curhaspath = self.haspath(matrix, rows, cols, path, x - 1, y, visited, pathlength) \
    #                      or self.haspath(matrix,
    #                                      rows,
    #                                      cols,
    #                                      path,
    #                                      x, y - 1,
    #                                      visited,
    #                                      pathlength) \
    #                      or self.haspath(
    #             matrix, rows, cols, path, x + 1, y, visited, pathlength) \
    #                      or self.haspath(matrix, rows, cols, path,
    #                                      x, y + 1, visited,
    #                                      pathlength)
    #         # 如果不能再走下一步，需要回退到上一状态
    #         if not curhaspath:
    #             pathlength[0] -= 1
    #             visited[y * cols + x] = False
    #     return curhaspath


if __name__ == "__main__":
    matrix = [
        ['a', 'b', 't', 'g'],
        ['c', 'f', 'c', 's'],
        ['j', 'd', 'e', 'h']
    ]
    # matrix = "abtgcfcsjdeh"
    rows, cols = 3, 4
    path = 'bfcedj'
    s = Solution()
    ret = s.hasPath(matrix, rows, cols, path)
    print(ret)
