import sys


class Solution:
    def FindTimes(self, m, n, x):
        ret = 0
        for i in range(m, n + 1):
            ret += self.CountNumber(i, x)
        return ret

    def CountNumber(self, n, x):
        times = 0
        while n:
            t = n % 10
            if t == x:
                times += 1
            n = n // 10
        return times


if __name__ == "__main__":
    readline = sys.stdin
    s = Solution()
    for line in readline:
        line = [int(i) for i in line.split(" ")]
        m, n, x = line[0], line[1], line[2]
        ret = s.FindTimes(m, n, x)
        print(ret)