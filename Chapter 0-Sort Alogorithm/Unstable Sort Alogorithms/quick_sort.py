from __future__ import print_function


def quick_sort(array):
    array_len = len(array)
    if array_len <= 1:
        return array
    else:
        pivot = array[0]
        greater = [element for element in array[1:] if element > pivot]
        lesser = [element for element in array[1:] if element <= pivot]
        return quick_sort(lesser) + [pivot] + quick_sort(greater)


if __name__ == '__main__':
    unsorted = [int(item) for item in input().split()]
    ret = quick_sort(unsorted)
    print(ret)