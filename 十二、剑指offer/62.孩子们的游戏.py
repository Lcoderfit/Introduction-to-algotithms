"""
62.孩子们的游戏.py
时间复杂度：O(n)
空间复杂度：O(1)
约瑟夫环公式：f(n, m) = [f(n-1, m) + m] % n
"""
# -*- coding:utf-8 -*-
class Solution:
    def LastRemaining_Solution(self, n, m):
        if n < 1 or m < 1:
            return -1
        last = 0
        for i in range(2, n + 1):
            last = (last + m) % i
        return last


if __name__ == "__main__":
    n = 5
    m = 3
    s = Solution()
    ret = s.LastRemaining_Solution(n, m)
    print(ret)