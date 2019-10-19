import sys


class Solution:
    def MathSum(self, array, sum):
        length = len(array)
        if length <= 1:
            return (-1, -1)
        array = sorted(array)
        i, j = 0, length - 1
        ret = (-1, -1)
        while i < j:
            if array[i] + array[j] == sum:
                ret = (i, j)
                break
            elif array[i] + array[j] > sum:
                j -= 1
            else:
                i += 1
        return ret


if __name__ == "__main__":
    s = Solution()
    n = int(input())
    array = [int(i) for i in input().split(" ")]
    sum = int(input())
    if n <= 0:
        print(-1, -1)
    else:
        ret = s.MathSum(array, sum)
        print(ret[0], ret[1])


