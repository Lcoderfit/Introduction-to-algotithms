"""

方法1： 位运算 （如果a^b == c, 则 a^c == b）
时间复杂度：O(n)
空间复杂度：O(n)

示例 1：

输入：encoded = [1,2,3], first = 1
输出：[1,0,2,1]
解释：若 arr = [1,0,2,1] ，那么 first = 1 且 encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [1,2,3]
示例 2：

输入：encoded = [6,2,7,3], first = 4
输出：[4,2,0,7,4]

"""
import sys
from typing import List


class Solution:
    @staticmethod
    def decode(encoded: List[int], first: int) -> List[int]:
        ans = [first]
        for v in encoded:
            t = v ^ first
            ans.append(t)
            first = t
        return ans


if __name__ == '__main__':
    pass
