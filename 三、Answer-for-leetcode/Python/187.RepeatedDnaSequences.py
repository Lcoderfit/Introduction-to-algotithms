from typing import List


class Solution1:
    '''用两个set进行处理'''
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        seen = set()
        res = set()
        for i in range(len(s)-9):
            substring = s[i:i+10]
            if substring in seen:
                res.add(substring)
            else:
                seen.append(substring)
        return list(res)


class Solution2:
    '''用一个dict和list处理'''
    def find_repeated_dna_sequences(self, s: str) -> List[str]:
        res = {}
        for i in range(len(s)-9):
            tmp = s[i:i+10]
            res[tmp] = res.get(tmp, 0) + 1
        return [i for i in res.keys() if res[i]>=2]


if __name__ == '__main__':
    while True:
        s = input()
        slt = Solution()
        res = slt.findRepeatedDnaSequences(s)
        print(res)