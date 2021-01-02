"""

方法1： 位运算
时间复杂度：O(1)
空间复杂度：O(1)

方法2： 字符串转换
时间复杂度：O(1)
空间复杂度：O(1)

case1:
10
r1:
5
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def find_complement(num: int) -> int:
        # 求数字a的补数b，则a的最左边的1对应b的位置，一定为0， 例如1010，最右边的1为第四位，其补数为0101，第四位变为0
        # 则b对应的10进制大小，首先求最左位1对应的满数位1的二进制数，例如a为1010，则先取2^4,，即10000
        # 然后将其减去1，得到1111，再用1111减去1010即求得b
        return 2 ** (len(bin(num)) - 2) - 1 - num

    @staticmethod
    def find_complement2(num: int) -> int:
        # int函数第二个参数表示第一个参数的进制数，第二个参数为2表示第一个参数为一个2进制数
        return int(bin(num)[2:].replace("0", "2").replace("1", "0").replace("2", "1"), 2)


if __name__ == '__main__':
    s = Solution()
    for line in sys.stdin:
        num_cur = int(line)
        # print(s.find_complement(num_cur))
        print(s.find_complement2(num_cur))
