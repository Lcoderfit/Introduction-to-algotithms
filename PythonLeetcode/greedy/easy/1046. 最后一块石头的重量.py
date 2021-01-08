"""

方法1： 大顶堆+模拟
时间复杂度：O(nlogn)
空间复杂度：O(n)

case1:
2 7 4 1 8 1
r:
1

Python中的切片操作会创建新的列表
"""
import heapq
import sys
from typing import List


class Solution:
    @staticmethod
    def last_stone_weight(stones: List[int]) -> int:
        heap = [-stone for stone in stones]
        # 构建最小堆
        heapq.heapify(heap)

        # Python只能构建小顶堆，所以通过取元素的相反数，构建小顶堆，则堆顶元素再取相反数后就是最大元素
        while len(heap) > 1:
            # 去最大的两个元素的相反数
            x, y = heapq.heappop(heap), heapq.heappop(heap)
            if x != y:
                # 设最大元素和第二大元素为a和b，则x=-a, y=-b, 要求a-b的差值的相反数再放入堆中，即求-(a-b) => -a-(-b) == x-y
                # 也可以这么想，x为最小元素，y为第二小元素，所以x - y求出的是其绝对值的相反数，也就是a与b绝对值的相反数
                heapq.heappush(heap, x - y)

        return -heap[0] if heap else 0


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        stones_cur = [int(i) for i in line.strip().split(" ")]
        print(s.last_stone_weight(stones_cur))
