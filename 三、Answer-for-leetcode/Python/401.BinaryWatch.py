from typing import List


class Solution:
    def read_binary_watch(self, num: int) -> List[str]:
        return ['%d:%02d' % (h, m)
                for h in range(12) for m in range(60)
                if (bin(h) + bin(m)).count('1') == num]


if __name__ == '__main__':
    while True:
        num = int(input())
        s = Solution()
        res = s.read_binary_watch(num)
        print(res)
