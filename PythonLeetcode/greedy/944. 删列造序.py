"""

方法1： 
时间复杂度：O()
空间复杂度：O()

case1:
cba daf ghi
r:
1

case1:
a b
r:
0
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def min_deletion_size(a: List[str]) -> int:
        if len(a[0]) == 1:
            return 0
        ans = 0
        for i in range(len(a[0])):
            for j in range(1, len(a)):
                if a[j][i] < a[j - 1][i]:
                    ans += 1
                    break
        return ans

    @staticmethod
    def min_deletion_size2(a: List[int]) -> int:
        length = 0
        # zip("ab", "de")会返回一个迭代器，迭代后元素依次为：('a', 'd'), ('b', 'e')
        for i in zip(*a):
            if list(i) != sorted(i):
                length += 1
        return length


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        a_cur = [i for i in line.strip().split(" ")]
        # print(s.min_deletion_size(a_cur))
        print(s.min_deletion_size2(a_cur))
