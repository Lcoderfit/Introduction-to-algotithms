"""
57-2.和为S的连续正数序列.py
时间复杂度：O(n)
空间复杂度：O(n)
"""
class Solution:
    def FindContinuousSequence(self, tsum):
        if tsum < 3:
            return list()
        small, big = 1, 2
        curSum = small + big
        mid = (tsum + 1) // 2
        ret = list()

        while small < mid:
            if curSum == tsum:
                ret.append([i for i in range(small, big + 1)])
            while curSum > tsum and small < mid:
                curSum -= small
                small += 1
                if curSum == tsum:
                    ret.append([i for i in range(small, big + 1)])
            big += 1
            curSum += big
        return ret


if __name__ == "__main__":
    tsum = 15
    s = Solution()
    ret = s.FindContinuousSequence(tsum)
    print(ret)
