"""

方法1： 贪心
时间复杂度：O(nlogn)
空间复杂度：O(1)

case1:
1 3
2 2
3 1

4
r:
8
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def maximum_unit(box_types: List[List[int]], truck_size: int) -> int:
        box_types.sort(key=lambda x: x[1], reverse=True)
        ans = 0
        for i in range(len(box_types)):
            if truck_size < box_types[i][0]:
                ans += truck_size * box_types[i][1]
                break
            else:
                ans += box_types[i][0] * box_types[i][1]
                truck_size -= box_types[i][0]
        return ans


if __name__ == '__main__':
    s = Solution()
    box_types_cur = []
    while True:
        in_list = [int(i) for i in input().strip().split(" ")]
        if len(in_list) == 2:
            box_types_cur.append(in_list)
        elif len(in_list) == 1:
            truck_size_cur = in_list[0]
            break
        else:
            raise Exception("Input Error")
    print(box_types_cur)
    print(s.maximum_unit(box_types_cur, truck_size_cur))
