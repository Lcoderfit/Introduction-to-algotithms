"""

方法1： 贪心算法
时间复杂度：O(n)
空间复杂度：O(1)

case1:
1 2 3
r:
1

case1:
2 2 2 3 3
r:
2

注意：位运算性能一般比其他算术运算符要更好

由点到面分析，先考虑只有两个的情况，再考虑三个。。。。
可以假设只有两个元素的状态为最终状态的前一个状态
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def min_cost_to_move_chips(position: List[int]) -> int:
        odd = 0
        even = 0
        for i in position:
            if i & 1:
                odd += 1
            else:
                even += 1
        return min(odd, even)


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        position_cur = [int(i) for i in line.strip().split(" ")]
        print(s.min_cost_to_move_chips(position_cur))
