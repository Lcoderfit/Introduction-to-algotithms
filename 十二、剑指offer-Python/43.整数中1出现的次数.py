"""
-5.最小的K个数
时间复杂度：O(nlogk)
空间复杂度：O(k)
"""
class Solution:
    def NumberOf1Between1AndN_Solution(self, n):
        if n < 1:
            return 0
        ret = 0
        for i in range(1, n + 1):
            ret = ret + self.CountOne(i)
        return ret

    def CountOne(self, number):
        count = 0
        while number:
            if number % 10 == 1:
                count += 1
            number = number // 10
        return count


if __name__ == "__main__":
    n = 13

    s = Solution()
    ret = s.NumberOf1Between1AndN_Solution(n)
    print(ret)
