import sys


class Solution:
    def EatFish(self, n, m):
        n = sorted(n)
        for k in range(m):
            j = 1
            while n[0] == n[j]:
                j += 1
            n[j] += n[0]
            del n[0]
            n.sort()
            print(n)
        return n[0]


if __name__ == "__main__":
    readline = sys.stdin
    s = Solution()
    for line in readline:
        line = [int(i) for i in line.split(" ")]
        n = line[2:]
        m = line[1]
        ret = s.EatFish(n, m)
        print(ret)

