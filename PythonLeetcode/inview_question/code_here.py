import math
import sys
from typing import List


class Solution1:
    @staticmethod
    def count_negative(nums: List[int]) -> (int, int):
        count_negative = 0
        ans = 0
        for i in range(len(nums)):
            if nums[i] < 0:
                count_negative += 1
            else:
                ans += nums[i]
        if (len(nums) - count_negative) == 0:
            return count_negative, None
        return count_negative, round(ans / (len(nums) - count_negative), 1)


class Solution2:
    @staticmethod
    def is_power_of_two(n: int) -> bool:
        return (n > 1) and ((n & (n - 1)) == 0)


class Solution3:
    @staticmethod
    def is_prime_number(n: int) -> bool:
        if n == 1:
            return False
        border = math.floor(math.sqrt(n)) + 1
        for i in range(2, border):
            if n % i == 0:
                return False
        return True


if __name__ == '__main__':
    # s = Solution1()
    # for line in sys.stdin:
    #     list_cur = [int(i) for i in line.split(" ")]
    #     num, nums_cur = list_cur[0], list_cur[1:]
    #     res = s.count_negative(nums_cur)
    #     print(*res)

    # s = Solution2()
    # for line in sys.stdin:
    #     n_cur = int(line.strip())
    #     print(s.is_power_of_two(n_cur))

    s = Solution3()
    for line in sys.stdin:
        n_cur = int(line.strip())
        print(s.is_prime_number(n_cur))
