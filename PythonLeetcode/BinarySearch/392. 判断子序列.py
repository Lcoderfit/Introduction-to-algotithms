"""
392. 判断子序列.py
时间复杂度：O(m+n)
空间复杂度：O(1)
"""
import sys


# A1: 贪心算法
class Solution:
    @staticmethod
    def is_sub_sequence(s: str, t: str) -> bool:
        i, j = 0, 0
        while (i < len(s)) and (j < len(t)):
            if s[i] == t[j]:
                i += 1
            j += 1
        if i == len(s):
            return True
        return False


# A2: 动态规划


if __name__ == '__main__':
    sl = Solution()
    for line in sys.stdin:
        line = [si for si in line.split(" ")]
        s_cur, t_cur = line[0], line[1]
        print(sl.is_sub_sequence(s_cur, t_cur))
