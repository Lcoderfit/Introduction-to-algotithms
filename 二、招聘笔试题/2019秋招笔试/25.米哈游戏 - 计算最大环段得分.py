import sys
from typing import List


class Solution:
    def maxScores(self, array: List[int]):
        if not array:
            return 0
        length = len(array)
        if length == 1:
            return array[0]
        ret = array[0]
        sum = array[0]
        i = 1
        k = 0
        while k < length and ((i < length and i > k) or (i >= length and (i % length < k))):
            t = i % length
            sum += array[t]
            if sum < 0:
                sum = array[t]
                k = i
            i += 1
            if ret < sum:
                ret = sum
        return ret


if __name__ == "__main__":
    t = input()
    array = [int(i) for i in input().split(" ")]
    s = Solution()
    ret = s.maxScores(array)
    print(ret)
