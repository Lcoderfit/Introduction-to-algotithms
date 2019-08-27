from typing import List


class Solution:
    def transfor(self, string: str) -> str:
        str_list = string.split(" ")
        num_str = list()
        operator_str = list()
        for i in str_list:
            if i.isdigit() or i[0] == '-':
                num_str.append(i)
            else:
                operator_str.append(i)

        print(num_str)
        print(operator_str)

        length = len(operator_str)
        if length == 1:
            s = " " + operator_str[0] + " "
            return s.join(sorted(num_str))

        t = operator_str[0]
        pos = 0
        ret = ""
        for i in range(1, length):
            if (t != operator_str[i]) or (i == length - 1):
                splt = " " + t + " "

                ret = ret + splt.join(sorted(num_str[pos:i])) + splt
                print(ret)
                t = operator_str[i]
                print("t：%s" % t)
                pos = i
        ret = ret + num_str[-1]
        return ret

if __name__ == "__main__":
    string = input().strip()
    s = Solution()
    ret = s.transfor(string)
    print(ret)