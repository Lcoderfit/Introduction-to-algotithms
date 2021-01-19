"""

方法1： 
时间复杂度：O(n)
空间复杂度：O(n)

case1:
5 2 C D +
r:
30

case2:
5 -2 4 C D 9 + +
r:

case3:
1
r:
1
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def cal_points(ops: List[str]) -> int:
        stack = []
        for op in ops:
            if op != "+" and op != "C" and op != "D":
                stack.append(int(op))
            elif stack:
                if (op == "+") and len(stack) >= 2:
                    top = stack.pop()
                    stack.extend([top, top + stack[-1]])
                elif op == "D":
                    top = stack.pop()
                    stack.extend([top, top * 2])
                elif op == "C":
                    stack.pop()
        print(stack)
        return sum(stack)


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        ops_cur = [i for i in line.strip().split(" ")]
        print(s.cal_points(ops_cur))
