"""
13.机器人的运动范围
时间复杂度：O(???)
空间复杂度：O(mn)
"""
class Solution:
    def movingCount(self, threshold, rows, cols):
        if threshold <0 and rows <= 0 and cols <= 0:
            return 0
        visited = [[0]*cols for _ in range(rows)]
        count = self.movingCountCore(threshold, rows, cols, 0, 0, visited)
        return count

    def movingCountCore(self, threshold, rows, cols, row, col, visited):
        count = 0
        if self.check(threshold, rows, cols, row, col, visited):
            visited[row][col] = 1
            count = 1 + self.movingCountCore(threshold, rows, cols, row - 1, col, visited) \
                    + self.movingCountCore(threshold, rows, cols, row, col + 1, visited) \
                    + self.movingCountCore(threshold, rows, cols, row + 1, col, visited) \
                    + self.movingCountCore(threshold, rows, cols, row, col - 1, visited)
        return count

    def check(self, threshold, rows, cols, row, col, visited):
        if (row >= 0 and row < rows) and (col >= 0 and col < cols) and (not visited[row][col]) \
            and (self.getDigitSum(row) + self.getDigitSum(col) <= threshold):
            return True
        return False



    def getDigitSum(self, number):
        sum = 0
        while number:
            sum += number % 10
            number = number // 10

        return sum


if __name__ == "__main__":
    rows, cols = 4, 4
    threshold = 3
    s = Solution()
    ret = s.movingCount(threshold, rows, cols)
    print(ret)