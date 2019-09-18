#coding=utf-8
import sys

class Solution:
    def merge_sort(self, collection):
        length = len(collection)
        if length > 1:
            midpoint = length // 2
            left_half = self.merge_sort(collection[:midpoint])
            right_half = self.merge_sort(collection[midpoint:])
            i, j, k = 0, 0, 0
            left_length = len(left_half)
            right_length = len(right_half)

            while i < left_length and j < right_length:
                if left_half[i] < right_half[j]:
                    collection[k] = left_half[i]
                    i += 1
                else:
                    collection[k] = right_half[j]
                    j += 1

                k += 1

            while i < left_length:
                collection[k] = left_half[i]
                i += 1
                k += 1

            while j < right_length:
                collection[k] = right_half[j]
                j += 1
                k += 1

        return collection

if __name__ == "__main__":
    str = sys.stdin.readline().strip()
    result = ""
    # 请在这里输入你的代码
    str_list = str.split(",")
    s = Solution()
    result = ",".join(s.merge_sort(str_list))
    print(result)