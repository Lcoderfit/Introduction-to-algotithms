"""

方法1： 
时间复杂度：O()
空间复杂度：O()

case1:
4 3 4 1
5 3 3 4 1
r:
3

case2:
1 2 2 3 4
3 4 1 2
r:
3

case3:
1 2 3
r:
1 2 3 4

case4:
4 5 6
3 3 3 3 3
r:
0
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def max_boxes_in_warehouse(boxes: List[int], warehouse: List[int]) -> int:
        index_warehouse = [(i, block) for i, block in enumerate(warehouse)]
        index_warehouse.sort(key=lambda x: (x[1], x[0]))
        hash_table = {}
        right_index = len(index_warehouse)

        for i, block in index_warehouse:
            if i >= right_index:
                continue
            if block not in hash_table:
                hash_table[block] = right_index - i
            else:
                hash_table[block] += right_index - i
            big_index = i

        print(hash_table)

        ans = 0
        warehouse.sort()
        i = 0
        for k, v in hash_table.items():
            while k >= warehouse[i] and v > 0:
                ans += 1
                v -= 1
                i += 1
        return ans


if __name__ == '__main__':
    s = Solution()
    boxes_cur = [int(i) for i in input().strip().split(" ")]
    warehouse_cur = [int(i) for i in input().strip().split(" ")]
    print(s.max_boxes_in_warehouse(boxes_cur, warehouse_cur))
