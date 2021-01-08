"""

方法1： 
时间复杂度：O(n)
空间复杂度：O(1)

case1:
ababcbacadefegdehijhklij
r:
[9, 7, 8]

case2:
abababfababasjdkfkjkjkjasiuyiqweuriansv
r:
[38, 1]
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def partition_labels(s: str) -> List[int]:
        hash_table = {}
        for i in range(len(s)):
            hash_table[s[i]] = i
        i = 0
        ans = []
        while i < len(s):
            right = hash_table[s[i]]
            j = i
            while j <= right:
                if right < hash_table[s[j]]:
                    right = hash_table[s[j]]
                j += 1
            ans.append(j - i)
            i = j
        return ans

    @staticmethod
    def partition_labels2(s: str) -> List[int]:
        hash_table = [0] * 26
        # enumerate可以同时获取元素和索引
        for i, ch in enumerate(s):
            hash_table[ord(ch) - ord('a')] = i

        start, end = 0, 0
        ans = []
        for i, ch in enumerate(s):
            end = max(end, hash_table[ord(s[i]) - ord('a')])
            if i == end:
                ans.append(end - start + 1)
                start = end + 1
        return ans


if __name__ == '__main__':
    sl = Solution()
    for line in sys.stdin:
        s_cur = line.strip()
        # print(sl.partition_labels(s_cur))
        print(sl.partition_labels2(s_cur))
