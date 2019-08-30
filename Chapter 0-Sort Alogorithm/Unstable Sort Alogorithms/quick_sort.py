"""
快速排序：不稳定排序，不稳定发生在中枢元素与其他元素交换的时候，5 3 3 4 3 8 9，中枢元素5与3交换会破坏稳定性
先分大小后递归，大于等于哨兵元素的放左边，小于哨兵元素的放右边

算法时间复杂度
1、最优：O(nlogn); 每一次取到的元素刚好平分整个数组，每一层时间n, 设层为k，2^k< n < 2^(k + 1)，k< logn < k + 1, 所以k取logn
2、平均：O(nlogn)
3、最差：每一次取到的元素刚好是最小的，O(n^2); T(n) = (n-1) + (n-1) + ... + 1 = n(n-1)/2 = O(n^2)
"""
from typing import List


def quick_sort1(array):
    array_len = len(array)
    if array_len <= 1:
        return array
    else:
        pivot = array[0]
        greater = [element for element in array[1:] if element > pivot]
        lesser = [element for element in array[1:] if element <= pivot]
        return quick_sort1(lesser) + [pivot] + quick_sort1(greater)


def partition(array, low, high):
    j = low - 1
    pivot = array[high]

    for i in range(low, high):
        if array[i] <= pivot:
            j += 1
            array[j], array[i] = array[i], array[j]

    array[j + 1], array[high] = array[high], array[j + 1]
    return j + 1

def quick_sort2(array, low, high):
    if low < high:
        q = partition(array, low, high)
        quick_sort2(array, low, q - 1)
        quick_sort2(array, q + 1, high)

    return array

import random


def rand_for_array(n: int, a: int, b: int) -> List:
    if n > abs(b - a + 1):
        return []

    array = list()
    arr_set = set()
    for i in range(n):
        while True:
            number = random.randint(a, b)
            if number not in arr_set:
                arr_set.add(number)
                break

        array.append(number)

    return array


if __name__ == '__main__':
    # unsorted1 = [int(item) for item in input().split()]
    # unsorted2 = unsorted1.copy()
    #
    # ret1 = quick_sort1(unsorted1)
    # print(ret1)
    #
    # ret2 = quick_sort2(unsorted2, 0, len(unsorted2) - 1)
    # print(ret2)
    array = rand_for_array(5, 1, 10)

    ret1 = quick_sort1(array)
    print(ret1)

    ret2 = quick_sort2(array, 0, len(array) - 1)
    print(ret2)