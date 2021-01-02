"""
类似于“含有重复元素的旋转数组”系列的题目，当遇到了特殊值导致正常的二分无法继续时，就临时退化为线性遍历
方法1： 二分查找 + 线性探测
时间复杂度：O(logn)
空间复杂度：O(1)

case1:
"ta"
"at" "" "" "" "ball" "" "" "car" "" "" "dad" "" ""

case2:
"ball"
"at" "" "" "" "ball" "" "" "car" "" "" "dad" "" ""
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def find_string(words: List[str], s: str) -> int:
        i, j = 0, len(words) - 1
        while i <= j:
            mid = (i + j) // 2
            tmp = mid
            while mid <= j and words[mid] == '':
                mid += 1
            if mid > j:
                j = tmp - 1
                continue
            if words[mid] == s:
                return mid
            elif words[mid] > s:
                j = tmp - 1
            else:
                i = mid + 1
        return -1


if __name__ == '__main__':
    sl = Solution()
    s_cur = input().strip('"')
    words_cur = [i.strip('"') for i in input().split(" ")]
    print(sl.find_string(words_cur, s_cur))
