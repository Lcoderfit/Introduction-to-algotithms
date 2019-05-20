from typing import List


class Solution:
    def candy(self, ratings: List[int]) -> int:
        s = 0
        n = len(ratings)
        s += n

        tmp = [0] * n
        for i in range(1, n):
            if ratings[i] > ratings[i-1]:
                tmp[i] = tmp[i-1] + 1
        for i in range(n-2, -1, -1):
            if ratings[i] > ratings[i+1]:
                tmp[i] = max(tmp[i], tmp[i+1]+1)
        s += sum(tmp)
        return s


if __name__ == '__main__':
    while True:
        ratings = list(map(int, input().split()))

        s = Solution()
        res = s.candy(ratings)
        print(res)