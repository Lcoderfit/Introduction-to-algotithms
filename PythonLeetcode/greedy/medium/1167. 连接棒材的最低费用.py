"""

方法1： 小根堆(优先队列) + 贪心
时间复杂度：O(nlog(n))
空间复杂度：O(n)

case1:
2 4 3
r:
14

case2:
1 8 3 5
r:
30

注意：遇到这种需要排序，然后取最小的两个或者最大的两个数进行操作得到一个结果，这个结果又要与剩下的元素进行同样操作的时候，可以采用堆的数据结构简化
"""
import heapq
import sys
from typing import List


class Solution:
    # sticks：木材
    @staticmethod
    def connect_sticks(sticks: List[int]) -> int:
        heapq.heapify(sticks)
        ans = 0
        while len(sticks) > 1:
            x, y = heapq.heappop(sticks), heapq.heappop(sticks)
            ans += x + y
            heapq.heappush(sticks, x + y)
        return ans


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        sticks_cur = [int(i) for i in line.strip().split(" ")]
        print(s.connect_sticks(sticks_cur))
