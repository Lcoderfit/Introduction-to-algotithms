class Solution:
    def max_hight(self, n):
        if n <= 0:
            return 0
        dp = [0] * (n + 1)
        dp[1] = 3
        for i in range(2, n + 1):
            dp[i] = dp[i - 1] + dp[i - 2]
        ret = 0
        for i in range(n, 0, -1):
            ret = (ret + dp[i]) * 2
        return ret


if __name__ == "__main__":
    n = int(input().strip())
    s = Solution()
    ret = s.max_hight(n)
    print(ret)