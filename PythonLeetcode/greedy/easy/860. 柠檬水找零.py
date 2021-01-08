"""

方法1： 贪心算法
时间复杂度：O(n)
空间复杂度：O(1)

case1:
5 5 5 10 20
r:
True

case2：
5 5 10
r:
True

case3:
10 10
r:
False

case4:
5 5 10 10 20
r:
False
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def lemonade_change(bills: List[int]) -> bool:
        five_count, ten_count = 0, 0
        for i in bills:
            if i == 5:
                five_count += 1
            elif i == 10:
                if five_count > 0:
                    ten_count += 1
                    five_count -= 1
                else:
                    return False
            else:
                if ten_count > 0 and five_count > 0:
                    ten_count -= 1
                    five_count -= 1
                elif five_count >= 3:
                    five_count -= 3
                else:
                    return False
        return True


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        bills_cur = [int(i) for i in line.strip().split(" ")]
        print(s.lemonade_change(bills_cur))
