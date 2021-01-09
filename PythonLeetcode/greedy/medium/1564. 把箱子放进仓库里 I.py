"""

方法1： hash表+排序+贪心
时间复杂度：O(wlogw+blob)
空间复杂度：O(b)

方法2： 排序+贪心
时间复杂度：O(w+blogb)
空间复杂度：O(1)


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
1 2 3 4
r:
1

case4:
4 5 6
3 3 3 3 3
r:
0

case4:
4 4 1 1
5 4 3 3 1
r:
4

1.当有双指针时，其中一个指针必须遍历完所有元素，则可将while替换为for循环
2.这个跟分发饼干有些类似，关键在于将高低不同的warehouse转换为非递增的序列
例如 3 5 4 2 3， 后一块墙能通过多大的板子受到前一块板子的限制，也就是能通过当前墙面的最大板子为min(前一块墙高度，当前墙高度)
所以可以通过从左到右遍历，两个相邻的墙对比，如果后一块墙要高于前一快，则将后一块改成跟前一块一样高就行,
再对boxes进行排序，这样warehouse和boxes都是有序的，两个有序序列的分发问题，就是分发饼干了，用双指针即可。
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
            right_index = i

        ans = 0
        boxes.sort()
        i = 0
        for k, v in hash_table.items():
            while (i < len(boxes)) and (k >= boxes[i]) and (v > 0):
                ans += 1
                v -= 1
                i += 1
        return ans

    @staticmethod
    def max_boxes_in_warehouse2(boxes: List[int], warehouse: List[int]) -> int:
        for i in range(1, len(warehouse)):
            if warehouse[i] > warehouse[i - 1]:
                warehouse[i] = warehouse[i - 1]
        warehouse.reverse()
        boxes.sort()
        i, j = 0, 0
        ans = 0
        while i < len(boxes) and j < len(warehouse):
            if boxes[i] <= warehouse[j]:
                i += 1
                ans += 1
            j += 1
        return ans

    @staticmethod
    def max_boxes_in_warehouse3(boxes: List[int], warehouse: List[int]) -> int:
        for i in range(1, len(warehouse)):
            if warehouse[i] > warehouse[i - 1]:
                warehouse[i] = warehouse[i - 1]
        boxes.sort()
        j, ans = 0, 0
        for i in range(len(warehouse) - 1, -1, -1):
            if (j < len(boxes)) and (warehouse[i] >= boxes[j]):
                ans += 1
                j += 1

        return ans


if __name__ == '__main__':
    s = Solution()
    while True:
        boxes_cur = [int(i) for i in input().strip().split(" ")]
        warehouse_cur = [int(i) for i in input().strip().split(" ")]

        if (boxes_cur[0] == -1) or (warehouse_cur[0] == -1):
            break

        # print(s.max_boxes_in_warehouse(boxes_cur, warehouse_cur))
        # print(s.max_boxes_in_warehouse2(boxes_cur, warehouse_cur))
        print(s.max_boxes_in_warehouse3(boxes_cur, warehouse_cur))
