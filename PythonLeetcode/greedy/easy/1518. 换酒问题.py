"""

方法1： 模拟法
时间复杂度：O(logn)
空间复杂度：O(1)

方法2： 模拟法
时间复杂度：O(b/e) 设有b瓶酒，e个空瓶换一瓶酒
空间复杂度：O(1)

方法3： 数学法
时间复杂度：O(1)
空间复杂度：O(1)

case1:
9 3
r:
13

case2:
15 4
r:
19

case3:
5 5
r:
6

case4:
2 3
r:
2
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def num_water_bottles(num_bottles: int, num_exchange: int) -> int:
        bottles, ans = num_bottles, num_bottles
        while bottles >= num_exchange:
            num_bottles = bottles // num_exchange
            bottles = num_bottles + bottles % num_exchange
            ans += num_bottles
        return ans

    @staticmethod
    def num_water_bottles2(num_bottles: int, num_exchange: int) -> int:
        bottles, ans = num_bottles, num_bottles
        while bottles >= num_exchange:
            bottles -= num_exchange
            ans += 1
            bottles += 1
        return ans

    @staticmethod
    def num_water_bottles3(num_bottles: int, num_exchange: int) -> int:
        # e个空瓶可以换一瓶酒，一瓶酒喝完后会产生一个空瓶，也就是说每次换酒，失去了e个空瓶换回来一个空瓶，相当于失去了e-1个空瓶
        # 转换为求 b -n(e-1) < e满足条件的n的最小值，其中只有当一开始b >= e时才可以换酒
        # 最后求出 n= (b-e)//(e-1) + 1 (b >= e)
        return (num_bottles - num_exchange) // (num_exchange - 1) + 1 + num_bottles \
            if num_bottles >= num_exchange else num_bottles


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        num_bottles_cur, num_exchange_cur = [int(i) for i in line.strip().split(" ")]
        # print(s.num_water_bottles(num_bottles_cur, num_exchange_cur))
        # print(s.num_water_bottles2(num_bottles_cur, num_exchange_cur))
        print(s.num_water_bottles3(num_bottles_cur, num_exchange_cur))
