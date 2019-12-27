def shell_sort(items):
    n = len(items)
    step = n // 2
    while step > 0:
        for cur in range(step, n):
            i = cur
            while i >= step and items[i-step] > items[i]:
                items[i - step], items[i] = items[i], items[i-step]
                i -= step
        step = step // 2


if __name__ == "__main__":
    array = [9, 8, 4, 6, 2, 3, 1, 5, 7]
    shell_sort(array)
    print(array)