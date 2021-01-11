"""

方法1： 计数法
时间复杂度：O(n*(m+∣Σ∣))  n为列表长度，m为字符串平均长度，∣Σ∣为小写字母字符集长度（∣Σ∣==26）
空间复杂度：O(∣Σ∣)

case1:
bella label roller
r:
e l l

case2:
cool lock cook
r:
c o
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def common_chars(a: List[str]) -> List[str]:
        # freq：频率
        # 遍历每个字符串，取每个字母出现的最小值，例如A字符串中a字符出现10次，但是B字符串中只出现2次，那么实际上“a”在这两个字符串中共同出现的次数为2次
        # 获取整型最大值，浮点型最大值用 float("inf")
        min_freq = [sys.maxsize] * 26

        for i in range(len(a)):
            # 用于记录每个单词中每个字符出现的频率
            freq = [0] * 26
            for ch in a[i]:
                freq[ord(ch) - ord('a')] += 1
            for j in range(26):
                # 取每个字符在每个单词中出现的最小值，例如A字符串中‘x’出现1次，但是B中没有出现，取最小值也就是0次
                min_freq[j] = min(min_freq[j], freq[j])
        ans = []
        # 如果某个字母不是在所有单词中都出现的字母，则min_freq中对应该字母的统计结果则为0， [chr(i + ord('a'))] * 0 是空列表，
        # ans.extend([])，这样的操作不会对ans有影响
        for i in range(26):
            ans.extend([chr(i + ord('a'))] * min_freq[i])
        return ans


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        a_cur = [i for i in line.strip().split(" ")]
        print(s.common_chars(a_cur))
