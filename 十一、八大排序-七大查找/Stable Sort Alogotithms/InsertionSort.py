"""
插入排序：稳定排序，当a[index - 1] > a[index]时才交换，相等时不交互
算法时间复杂度
1、最优：O(n), 数组恰好有序，插入每个元素时都只比一次
2、平均：O(n^2)
3、最差：O(n^2), 第i个元素需要比较i - 1次，1 + 2 + ... + n - 1
"""
class Solution:
    def insertion_sort(self, collection):
        for index in range(len(collection)):
            while index > 0 and collection[index - 1] > collection[index]:
                collection[index - 1], collection[index] = collection[index], collection[index - 1]
                index -= 1

        return collection


if __name__  == "__main__":
    array = [int(item) for item in input().split()]
    s = Solution()
    ret = s.insertion_sort(array)
    print(ret)
