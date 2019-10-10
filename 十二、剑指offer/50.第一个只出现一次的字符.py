"""
50.第一个只出现一次的字符
时间复杂度：O(n)
空间复杂度：O(1)
"""
class Solution:
    def FirstNotRepeatingChar(self, s):
        if not s:
            return -1
        char_list = [0] * 256

        for i in s:
            char_list[ord(i)] += 1

        ret = -1
        for i in range(len(s)):
            if char_list[ord(s[i])] == 1:
                ret = i
                break

        return ret


if __name__ == "__main__":
    string = "google"

    s = Solution()
    ret = s.FirstNotRepeatingChar(string)
    print(ret)
