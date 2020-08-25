import sys


# 最大公约数
class Solution:
    def gcd1(self, x, y):
        while y != 0:
            x, y = y, x % y

        return x

    def gcd2(self, x, y):
        while y != 0:
            if x > y:
                x = x - y
            else:
                y = y - x

        return x


# 最小公倍数
class Solution1:
    def gcd1(self, x, y):
        s = x * y
        while y != 0:
            x, y = y, x % y

        return s // x

    def gcd2(self, x, y):
        s = x * y
        while y != 0:
            if x > y:
                x = x - y
            else:
                y = y - x

        return s // x


if __name__ == "__main__":
    # readline = sys.stdin
    # for line in readline:
    #     line = [int(i) for i in line.split(" ")]
    #     x, y = int(line[0]), int(line[1])
    #     s = Solution()
    #     res1 = s.gcd1(x, y)
    #     res2 = s.gcd2(x, y)
    #     print(res1, res2)
    #
    #     s1 = Solution1()
    #     res3 = s1.gcd1(x, y)
    #     res4 = s1.gcd2(x, y)
    #     print(res3, res4)
    record = ('Dave', 'dave#example.com', '777-555-430', '432-532-52')
    name, email, *phone = record
    print(phone)
