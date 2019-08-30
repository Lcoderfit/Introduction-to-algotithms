"""
冒泡排序：稳定排序，当a[j] > a[j + 1]时才交换，相等时不交互
算法时间复杂度
1、最优：冒泡排序优化一的最优情况是O(n)
2、平均：有n-1种情况，每种情况需要比较的次数为i*(n - 1)， (1+2+...+n-1)*(n-1)/(n-1) = O(n^2)
3、最差：O(n^2)
"""


# 冒泡排序优化一
class BubbleSort1:
    def bubble_sort(self, collection):
        length = len(collection)
        for i in range(length - 1):
            swapped = False
            for j in range(length - i - 1):
                if collection[j] > collection[j + 1]:
                    swapped = True
                    collection[j], collection[j + 1] = collection[j + 1], collection[j]
            if not swapped: break

        return collection

    @staticmethod
    def print_list(collection):
        for i in collection:
            print(i, end=" ")
        print()


# 冒泡排序优化二
class BubbleSort2:
    def bubble_sort(self, collection):
        length = len(collection)
        k = length - 1
        for i in range(length - 1):
            swapped = False
            pos = 0
            for j in range(k):
                if collection[j] > collection[j + 1]:
                    swapped = True
                    pos = j
                    collection[j], collection[j + 1] = collection[j + 1], collection[j]
            if not swapped: break
            k = pos
        return collection

    @staticmethod
    def print_list(collection):
        for i in collection:
            print(i, end=" ")
        print()


# 冒泡排序，优化方案三
class BubbleSort3:
    def bubble_sort(self, collection):
        length = len(collection)
        borden_right = length - 1
        borden_left = 0
        pos_right = 0
        pos_left = 0
        for i in range(length - 1):
            swapped = False
            for j in range(borden_left, borden_right):
                if collection[j] > collection[j + 1]:
                    swapped = True
                    pos_right = j
                    collection[j], collection[j + 1] = collection[j + 1], collection[j]

            if not swapped: break
            swapped = False
            borden_right = pos_right

            for j in range(borden_right, borden_left, -1):
                if collection[j] < collection[j - 1]:
                    swapped = True
                    pos_left = j
                    collection[j], collection[j - 1] = collection[j - 1], collection[j]

            if not swapped: break
            borden_left = pos_left

        return collection

    @staticmethod
    def print_list(collection):
        for i in collection:
            print(i, end=" ")
        print()

    def bubble_sort5(self, collection):
        length = len(collection)
        borden_left = 0
        borden_right = length - 1
        pos_left = 0
        pos_right = 0

        for i in range(0, length - 1):
            swapped = False
            for j in range(borden_left, borden_right):
                if collection[j] > collection[j + 1]:
                    swapped = True
                    pos_right = j
                    collection[j], collection[j + 1] = collection[j + 1], collection[j]

            if not swapped: break
            swapped = False
            borden_right = pos_right

            for j in range(borden_right, borden_left, -1):
                if collection[j] < collection[j - 1]:
                    swapped = True
                    pos_left = j
                    collection[j], collection[j - 1] = collection[j - 1], collection[j]

            if not swapped: break
            borden_left = pos_left

        return collection


if __name__ == "__main__":
    int_list = list(map(int, input().split()))
    bs = BubbleSort3()
    int_list = bs.bubble_sort(int_list)
    bs.print_list(int_list)
