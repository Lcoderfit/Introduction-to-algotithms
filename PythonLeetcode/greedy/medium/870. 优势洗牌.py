"""

方法1： 排序+贪心 （单指针插入法）
时间复杂度：O(nlogn)
空间复杂度：O(n)

方法2： 排序+贪心 （双表拆合法）
时间复杂度：O(nlogn)
空间复杂度：O(n)

case1:
2 7 11 15
1 10 4 11
r:
2 11 7 15

case2:
12 24 8 32
13 25 32 11
r:
24 32 8 12
"""
import sys
from typing import List


class Solution:
    # 用单指针位移来将元素插入到对应的结果列表里
    @staticmethod
    def advantage_count(a: List[int], b: List[int]) -> List[int]:
        a.sort()
        ans = [-1] * len(a)
        # 索引序列，用来记录元素排序后最原始的索引位置
        c = [(i, v) for i, v in enumerate(b)]
        c.sort(key=lambda x: (x[1]))

        j, k = 0, 0
        for i in range(len(a)):
            while ans[k] != -1:
                k += 1
            if a[i] > c[j][1]:
                if ans[c[j][0]] != -1:
                    ans[k] = ans[c[j][0]]
                ans[c[j][0]] = a[i]
                j += 1
            else:
                ans[k] = a[i]

        return ans

    @staticmethod
    def advantage_count2(a: List[int], b: List[int]) -> List[int]:
        sorted_a = sorted(a)
        sorted_b = sorted(b)

        # 指定的需要移动到对应位置的元素
        assigned = {v: [] for v in b}
        # 可以随便放的，只要不跟指定位置冲突的元素
        remaining = []

        j = 0
        for v in sorted_a:
            if v > sorted_b[j]:
                assigned[sorted_b[j]].append(v)
                j += 1
            else:
                remaining.append(v)

        # 元素弹出列表的顺序无关紧要，因为首先remaining中的元素不需要放置到特定索引位置，所以其中元素可以放到任意其他位置
        # 而assigned字典中，如果同一个key对应的value（列表）中包含多个元素，则其弹出顺序也无关紧要，因为都这些元素都要大于key（key对应b中元素）
        return [assigned[v].pop() if assigned[v] else remaining.pop() for v in b]


if __name__ == '__main__':
    s = Solution()
    while True:
        a_cur = [int(i) for i in input().strip().split(" ")]
        b_cur = [int(i) for i in input().strip().split(" ")]

        if (a_cur[0] == -1) or (b_cur[0] == -1) or (len(a_cur) != len(b_cur)):
            break
        # print(s.advantage_count(a_cur, b_cur))
        print(s.advantage_count2(a_cur, b_cur))
