"""
121.买股票的最佳时期：O(n)
temp - prices[i]小于0时，表示有比temp更加小的变量，所以则将temp更新为更小的变量
"""
class Solution:
    def max_Profile(self, prices):
        length = len(prices)
        if length <= 1:
            return 0
        temp = prices[0]
        max = 0
        for i in range(length):
            div = prices[i] - temp
            if div <= 0:
                temp = prices[i]
            else:
                if max < div:
                    max = div

        return max


if __name__ == "__main__":
    prices = [int(i) for i in input().split()]
    s = Solution()
    ret = s.max_Profile(prices)
    print(ret)