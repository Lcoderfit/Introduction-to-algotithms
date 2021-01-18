"""

方法1： 栈
时间复杂度：O(n)
空间复杂度：O(n)

方法2： 双指针
时间复杂度：O(n)
空间复杂度：O(1)

case1:
ab#c ad#c
r:
True

case2:
"ab##" "c#d#"
r:
True

case3:
a##c #a#c
r:
true

case4:
a#c b
r:
False

case5:
bbbextm bbb#extm
r:
False
"""
import sys
from typing import List


class Solution:
    def back_space_compare(self, s: str, t: str) -> bool:
        return self.back_space(s) == self.back_space(t)

    @staticmethod
    def back_space(s: str):
        stack = []
        for ch in s:
            # 判断条件尽量简化
            if ch != "#":
                stack.append(ch)
            elif stack:
                stack.pop()
        return "".join(stack)

    @staticmethod
    def back_space_compare2(s: str, t: str) -> bool:
        cnt1, cnt2 = 0, 0
        i, j = len(s) - 1, len(t) - 1
        # 因为可能会产生i与j其中一个大于等于0，另一个小于0的情况，用or的话就可以再次循环，
        # 在循环体内判断i与j是否一个大于0一个小于0，是则返回False
        while i >= 0 or j >= 0:
            while i >= 0:
                if s[i] == "#":
                    cnt1 += 1
                    i -= 1
                elif cnt1 > 0:
                    cnt1 -= 1
                    i -= 1
                else:
                    # 退出循环，用于比较s[i]和t[j]是否相等
                    break

            while j >= 0:
                if t[j] == "#":
                    cnt2 += 1
                    j -= 1
                elif cnt2 > 0:
                    cnt2 -= 1
                    j -= 1
                else:
                    break

            if i >= 0 and j >= 0:
                if s[i] != t[j]:
                    return False
            # 一个遍历完了，一个还遍历完，说明两个字符串经过退格处理后长度不一样，肯定不相等
            elif i >= 0 or j >= 0:
                return False

            i -= 1
            j -= 1
        return True


if __name__ == '__main__':
    sl = Solution()
    s_cur, t_cur = input().strip().split(" ")
    # print(sl.back_space_compare(s_cur, t_cur))
    print(sl.back_space_compare2(s_cur, t_cur))
