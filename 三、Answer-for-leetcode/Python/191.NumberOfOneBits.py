class Solution:
    def number_of_one_bits(self, n: int) -> int:
        return bin(n).count('1')


if __name__ == '__main__':
    while True:
        n = int(input())
        s = Solution()
        res = s.number_of_one_bits(n)
        print(res)