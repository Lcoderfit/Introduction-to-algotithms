'''
9.回文数：palindrome-number
leetcode链接：https://leetcode-cn.com/problems/palindrome-number/
时间复杂度：O(log10(n))
空间复杂度：O(1)
'''


class Solution:
    def isPalindrome(self, x):
        '''
        :type x: int
        :return type:bool
        '''
        if x < 0:
            return False
        else:
            # [::-1]效果为：将字符串反转
            y = str(x)[::-1]
            if y == str(x):
                return True
            else:
                return False


if __name__ == "__main__":
    while True:
        x = input()
        x = int(x)
        s = Solution()
        res = s.isPalindrome(x)
        print(res)
