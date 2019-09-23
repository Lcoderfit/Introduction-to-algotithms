"""
1.朴素算法：O(nm)
2.KMP: O(n + m)

关键：j = next[j - 1]
index = next[index - 1]

母串：a b a d
子串：a b a c
next:0 0 1 0
下次：  (a) b a c
当字串匹配到下标为3的位置时，d与c不匹配，则下一次匹配是从母串的下标3（d）、字串的下标1（b）位置开始

子串：a b c d a b c k
     0 0 0 0 1 2 3 0
当计算子串的next时，遇到d与k不匹配的情况，则可以把a b c d看成子串，a b c k看成母串，则index = next[index - 1]
"""
class Solution:
    def kmp_algorithm(self, string, substring):
        pnext = self.gen_pnext(substring)
        n = len(string)
        m = len(substring)
        i, j = 0, 0
        while i < n and j < m:
            if string[i] == substring[j]:
                i += 1
                j += 1
            elif j != 0:
                j = pnext[j - 1]
            else:
                # 此时j == 0, 母串继续从i的后一个位置比较
                i += 1
        if j == m:
            # 0 1 2 3 4 5 6 7
            # a b c d e f g
            #         0 1 2 3
            #         e f g
            # 最后匹配结束，i == 7 j == 3
            # i - j 则为i中字符”e“的下标索引位置
            return i - j
        else:
            return -1

    def gen_pnext(self, substring):
        index, m = 0, len(substring)
        pnext = [0] * m
        i = 1
        while i < m:
            if substring[i] == substring[index]:
                # 0 1 2 3
                # a b a c
                # a b a k
                # 0 0 1 0
                # next[2] = 0 + 1 = 1 ===> 从索引为1处开始比较
                pnext[i] = index + 1
                index += 1
                i += 1
            elif index != 0:
                index = pnext[index - 1]
            else:
                pnext[i] = 0
                i += 1
        return pnext


if __name__ == "__main__":
    string = 'abcxabcdabcdabcy'
    substring = 'abcdabcy'
    s = Solution()
    ret = s.kmp_algorithm(string, substring)
    print(ret)






